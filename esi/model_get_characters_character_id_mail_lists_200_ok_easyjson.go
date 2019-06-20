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

func easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLists200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMailLists200OkList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdMailLists200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMailLists200Ok
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
func easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdMailLists200OkList) {
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
func (v GetCharactersCharacterIdMailLists200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLists200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLists200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLists200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdMailLists200Ok) {
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
		case "mailing_list_id":
			out.MailingListId = int32(in.Int32())
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
func easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdMailLists200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MailingListId != 0 {
		const prefix string = ",\"mailing_list_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.MailingListId))
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
func (v GetCharactersCharacterIdMailLists200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMailLists200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1e2f2d49EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLists200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMailLists200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1e2f2d49DecodeGithubComAntihaxGoesiEsi1(l, v)
}
