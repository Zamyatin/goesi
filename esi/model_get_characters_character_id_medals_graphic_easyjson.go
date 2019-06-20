// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdMedalsGraphicList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMedalsGraphicList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdMedalsGraphicList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMedalsGraphic
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdMedalsGraphicList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMedalsGraphicList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMedalsGraphicList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMedalsGraphicList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMedalsGraphicList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdMedalsGraphic) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "color":
			out.Color = int32(in.Int32())
		case "graphic":
			out.Graphic = string(in.String())
		case "layer":
			out.Layer = int32(in.Int32())
		case "part":
			out.Part = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdMedalsGraphic) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Color != 0 {
		const prefix string = ",\"color\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Color))
	}
	if in.Graphic != "" {
		const prefix string = ",\"graphic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Graphic))
	}
	if in.Layer != 0 {
		const prefix string = ",\"layer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Layer))
	}
	if in.Part != 0 {
		const prefix string = ",\"part\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Part))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMedalsGraphic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMedalsGraphic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6d4c19fbEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMedalsGraphic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMedalsGraphic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6d4c19fbDecodeGithubComAntihaxGoesiEsi1(l, v)
}
