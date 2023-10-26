package sexpr

import (
	"reflect"
	"fmt"
	"io"
)

type Encoder struct {
	w io.Writer
	err error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc *Encoder) Encode(v interface{}) error {
	return encode(enc.w, reflect.ValueOf(v))
}

func encode(w io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		io.WriteString(w, "nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(w, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(w, "%q", v.String())
	case reflect.Ptr:
		fmt.Fprint(w, v.Elem())
	case reflect.Array, reflect.Slice:
		io.WriteString(w, "(")
		for i := 0; i < v.Len(); i++ {
			if !v.Index(i).IsZero() {
				if i > 0 {
					io.WriteString(w, " ")
				}
				if err := encode(w, v.Index(i)); err != nil {
					return err
				}
			}
		}
		io.WriteString(w, ")")
	case reflect.Struct:
		io.WriteString(w, "(")
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).IsZero() {
				if i > 0 {
					io.WriteString(w, " ")
				}
				fmt.Fprintf(w, "(%s ", v.Type().Field(i).Name)
				if err := encode(w, v.Field(i)); err != nil {
					return err
				}
				io.WriteString(w, ")")
			}
		}
		io.WriteString(w, ")")
	case reflect.Map:
		io.WriteString(w, "(")
		for i, key := range v.MapKeys() {
			if !v.MapIndex(key).IsZero() {
				if i > 0 {
					io.WriteString(w, " ")
				}
				io.WriteString(w, "(")
				if err := encode(w, key); err != nil {
					return err
				}
				io.WriteString(w, " ")
				if err := encode(w, v.MapIndex(key)); err != nil {
					return err
				}
				io.WriteString(w, ")")
			}
		}
		io.WriteString(w, ")")
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}


/*func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}*/
