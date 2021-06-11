package ecore

import "encoding/xml"

const (
	xmiURI        = "http://www.omg.org/XMI"
	xmiNS         = "xmi"
	versionAttrib = "version"
	uuidAttrib    = "uuid"
)

type XMIDecoder struct {
	*XMLDecoder
	xmiVersion string
}

func NewXMIDecoder(options map[string]interface{}) *XMIDecoder {
	l := new(XMIDecoder)
	l.XMLDecoder = NewXMLDecoder(options)
	l.notFeatures = append(l.notFeatures,
		xml.Name{Space: xmiURI, Local: typeAttrib},
		xml.Name{Space: xmiURI, Local: versionAttrib},
		xml.Name{Space: xmiURI, Local: uuidAttrib},
	)
	l.extendedMetaData = nil
	l.interfaces = l
	return l
}

func (l *XMIDecoder) GetXMIVersion() string {
	return l.xmiVersion
}

func (l *XMIDecoder) getXSIType() string {
	xsiType := l.XMLDecoder.getXSIType()
	if len(xsiType) == 0 && l.attributes != nil {
		return l.getAttributeValue(xmiURI, typeAttrib)
	}
	return xsiType
}

func (l *XMIDecoder) handleAttributes(object EObject) {
	version := l.getAttributeValue(xmiURI, versionAttrib)
	if len(version) > 0 {
		l.xmiVersion = version
	}
	l.XMLDecoder.handleAttributes(object)
}
