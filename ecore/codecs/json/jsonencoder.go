package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/karlseguin/jsonwriter"
	"github.com/masagroup/soft.go/ecore"
)

type JSONEncoder struct {
	resource        ecore.EResource
	w               *jsonwriter.Writer
	featureKinds    map[ecore.EStructuralFeature]jsonFeatureKind
	errorFn         func(diagnostic ecore.EDiagnostic)
	idAttributeName string
	keepDefaults    bool
}

func NewJSONEncoder(resource ecore.EResource, w io.Writer, options map[string]interface{}) *JSONEncoder {
	e := &JSONEncoder{
		w:            jsonwriter.New(w),
		resource:     resource,
		featureKinds: map[ecore.EStructuralFeature]jsonFeatureKind{},
	}
	if options != nil {
		e.idAttributeName, _ = options[JSON_OPTION_ID_ATTRIBUTE_NAME].(string)
	}
	return e
}

func (e *JSONEncoder) EncodeResource() {
	e.errorFn = func(diagnostic ecore.EDiagnostic) {
		e.resource.GetErrors().Add(diagnostic)
	}
	if contents := e.resource.GetContents(); !contents.Empty() {
		object := contents.Get(0).(ecore.EObject)
		e.encodeTopObject(object)
	}
}

func (e *JSONEncoder) EncodeObject(object ecore.EObject) (err error) {
	e.errorFn = func(diagnostic ecore.EDiagnostic) {
		if err == nil {
			err = diagnostic
		}
	}
	e.encodeTopObject(object)
	return
}

func (e *JSONEncoder) encodeTopObject(eObject ecore.EObject) {
	e.w.RootObject(func() {
		e.encodeObject(eObject)
	})
}

func (e *JSONEncoder) encodeObject(eObject ecore.EObject) {
	eClass := eObject.EClass()
	// class
	e.w.KeyString("eClass", e.getClassName(eClass))
	// id
	if idManager := e.resource.GetObjectIDManager(); len(e.idAttributeName) > 0 && idManager != nil {
		if id := idManager.GetID(eObject); id != nil {
			e.w.KeyValue(e.idAttributeName, fmt.Sprintf("%v", id))
		}
	}
	// features
	for itFeature := eClass.GetEAllStructuralFeatures().Iterator(); itFeature.HasNext(); {
		eFeature := itFeature.Next().(ecore.EStructuralFeature)
		e.encodeFeatureValue(eObject, eFeature)
	}
}

func (e *JSONEncoder) encodeObjectReference(eObject ecore.EObject) {
	e.w.KeyString("eClass", e.getClassName(eObject.EClass()))
	e.w.KeyString("eRef", e.getReference(eObject))
}

func jsonEscape(i string) string {
	w := &bytes.Buffer{}
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	if err := e.Encode(i); err != nil {
		panic(err)
	}
	// last \n
	b := w.Bytes()
	return string(b[:len(b)-1])
}

func (e *JSONEncoder) encodeFeatureValue(eObject ecore.EObject, eFeature ecore.EStructuralFeature) {
	if !e.shouldSaveFeature(eObject, eFeature) {
		return
	}

	// compute feature kind
	kind, ok := e.featureKinds[eFeature]
	if !ok {
		kind = getJSONCodecFeatureKind(eFeature)
		e.featureKinds[eFeature] = kind
	}

	value := eObject.EGetResolve(eFeature, false)
	switch kind {
	case jfkTransient:
	case jfkData:
		str, ok := e.getData(value, eFeature)
		if ok {
			// don't use e.w.KeyString because json escaping is bugged
			// in jsonencoder
			e.w.Key(eFeature.GetName())
			e.w.Raw([]byte(jsonEscape(str)))
		}
	case jfkDataList:
		l := value.(ecore.EList)
		e.w.Array(eFeature.GetName(), func() {
			for it := l.Iterator(); it.HasNext(); {
				str, ok := e.getData(it.Next(), eFeature)
				if ok {
					e.w.Value(str)
				}
			}
		})
	case jfkObject:
		e.w.Object(eFeature.GetName(), func() {
			e.encodeObject(value.(ecore.EObject))
		})
	case jfkObjectList:
		l := value.(ecore.EList)
		e.w.Array(eFeature.GetName(), func() {
			for it := l.Iterator(); it.HasNext(); {
				e.w.ArrayObject(func() {
					e.encodeObject(it.Next().(ecore.EObject))
				})
			}
		})
	case jfkObjectReference:
		e.w.Object(eFeature.GetName(), func() {
			e.encodeObjectReference(value.(ecore.EObject))
		})
	case jfkObjectReferenceList:
		l := value.(ecore.EList)
		e.w.Array(eFeature.GetName(), func() {
			for it := l.Iterator(); it.HasNext(); {
				e.w.ArrayObject(func() {
					e.encodeObjectReference(it.Next().(ecore.EObject))
				})
			}
		})
	}
}

func (e *JSONEncoder) shouldSaveFeature(o ecore.EObject, f ecore.EStructuralFeature) bool {
	return o.EIsSet(f) || (e.keepDefaults && f.GetDefaultValueLiteral() != "")
}

func (e *JSONEncoder) getClassName(eClass ecore.EClass) string {
	ePackage := eClass.GetEPackage()
	return ePackage.GetNsURI() + "#//" + eClass.GetName()
}

func (e *JSONEncoder) getData(value interface{}, f ecore.EStructuralFeature) (string, bool) {
	if value == nil {
		return "", false
	} else {
		d := f.GetEType().(ecore.EDataType)
		p := d.GetEPackage()
		f := p.GetEFactoryInstance()
		s := f.ConvertToString(d, value)
		return s, true
	}
}

func (e *JSONEncoder) getReference(eObject ecore.EObject) string {
	eInternal, _ := eObject.(ecore.EObjectInternal)
	if eInternal != nil {
		objectURI := eInternal.EProxyURI()
		if objectURI == nil {
			eOtherResource := eObject.EResource()
			if eOtherResource == nil {
				if e.resource != nil && e.resource.GetObjectIDManager() != nil && e.resource.GetObjectIDManager().GetID(eObject) != nil {
					objectURI = e.getResourceReference(e.resource, eObject)
				} else {
					return ""
				}
			} else {
				objectURI = e.getResourceReference(eOtherResource, eObject)
			}
		}
		objectURI = e.resource.GetURI().Relativize(objectURI)
		return objectURI.String()
	}
	return ""
}

func (e *JSONEncoder) getResourceReference(resource ecore.EResource, object ecore.EObject) *ecore.URI {
	return ecore.NewURIBuilder(resource.GetURI()).SetFragment(resource.GetURIFragment(object)).URI()
}
