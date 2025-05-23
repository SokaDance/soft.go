package sqlite

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/masagroup/soft.go/ecore"
)

const (
	SQL_OPTION_KEEP_DEFAULTS         = "KEEP_DEFAULTS"         // keep default values
	SQL_OPTION_CODEC_VERSION         = "CODEC_VERSION"         // codec version ( int64 )
	SQL_OPTION_OBJECT_ID             = "OBJECT_ID"             // object id column name ( string )
	SQL_OPTION_CONTAINER_ID          = "CONTAINER_ID"          // container id ( boolean ) if true, encode object's container and container feature id
	SQL_OPTION_SQL_ID_MANAGER        = "SQL_ID_MANAGER"        // SQL id manager
	SQL_OPTION_IN_MEMORY_DATABASE    = "IN_MEMORY_DATABASE"    // uses a memory database
	SQL_OPTION_DECODER_WITH_FEATURES = "DECODER_WITH_FEATURES" // decoder with object features
	SQL_OPTION_DECODER_WITH_OBJECTS  = "DECODER_WITH_OBJECTS"  // decoder with objects
	SQL_OPTION_DECODER_DB_PATH       = "DECODER_DB_PATH"
	SQL_OPTION_LOGGER                = "LOGGER"
)

type SQLCodec struct {
}

const sqlCodecVersion int64 = 1

func sqlTmpDB(prefix string) (string, error) {
	try := 0
	for {
		randBytes := make([]byte, 16)
		_, err := rand.Read(randBytes)
		if err != nil {
			return "", err
		}
		f := filepath.Join(os.TempDir(), prefix+"."+hex.EncodeToString(randBytes)+".sqlite")
		_, err = os.Stat(f)
		if os.IsExist(err) {
			if try++; try < 10000 {
				continue
			}
			return "", &fs.PathError{Op: "sqlTmpDB", Path: prefix, Err: fs.ErrExist}
		}
		return f, nil
	}
}

func (d *SQLCodec) NewEncoder(resource ecore.EResource, w io.Writer, options map[string]any) ecore.EEncoder {
	return NewSQLWriterEncoder(w, resource, options)
}
func (d *SQLCodec) NewDecoder(resource ecore.EResource, r io.Reader, options map[string]any) ecore.EDecoder {
	return NewSQLReaderDecoder(r, resource, options)
}

type SQLCodecIDManager interface {
	SetPackageID(ecore.EPackage, int64)
	SetObjectID(ecore.EObject, int64)
	SetClassID(ecore.EClass, int64)
	SetEnumLiteralID(ecore.EEnumLiteral, int64)
}
