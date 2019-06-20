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

func easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostUniverseIdsStationList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUniverseIdsStationList, 0, 2)
			} else {
				*out = PostUniverseIdsStationList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUniverseIdsStation
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
func easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostUniverseIdsStationList) {
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
func (v PostUniverseIdsStationList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseIdsStationList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseIdsStationList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseIdsStationList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostUniverseIdsStation) {
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
		case "id":
			out.Id = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
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
func easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostUniverseIdsStation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Id))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUniverseIdsStation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseIdsStation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD078dd6EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseIdsStation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseIdsStation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD078dd6DecodeGithubComAntihaxGoesiEsi1(l, v)
}
