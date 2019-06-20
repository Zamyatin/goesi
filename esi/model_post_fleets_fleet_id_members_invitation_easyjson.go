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

func easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostFleetsFleetIdMembersInvitationList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostFleetsFleetIdMembersInvitationList, 0, 1)
			} else {
				*out = PostFleetsFleetIdMembersInvitationList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostFleetsFleetIdMembersInvitation
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
func easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostFleetsFleetIdMembersInvitationList) {
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
func (v PostFleetsFleetIdMembersInvitationList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersInvitationList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersInvitationList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersInvitationList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostFleetsFleetIdMembersInvitation) {
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
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "role":
			out.Role = string(in.String())
		case "squad_id":
			out.SquadId = int64(in.Int64())
		case "wing_id":
			out.WingId = int64(in.Int64())
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
func easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostFleetsFleetIdMembersInvitation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CharacterId))
	}
	if in.Role != "" {
		const prefix string = ",\"role\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Role))
	}
	if in.SquadId != 0 {
		const prefix string = ",\"squad_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SquadId))
	}
	if in.WingId != 0 {
		const prefix string = ",\"wing_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WingId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostFleetsFleetIdMembersInvitation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFleetsFleetIdMembersInvitation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFf081cddEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFleetsFleetIdMembersInvitation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFleetsFleetIdMembersInvitation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFf081cddDecodeGithubComAntihaxGoesiEsi1(l, v)
}
