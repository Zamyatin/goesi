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

func easyjson9f888528DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdFwStatsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdFwStatsOkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdFwStatsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdFwStatsOk
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
func easyjson9f888528EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdFwStatsOkList) {
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
func (v GetCorporationsCorporationIdFwStatsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f888528EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdFwStatsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f888528EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdFwStatsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f888528DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdFwStatsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f888528DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson9f888528DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdFwStatsOk) {
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
		case "enlisted_on":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EnlistedOn).UnmarshalJSON(data))
			}
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "kills":
			easyjson9f888528DecodeGithubComAntihaxGoesiEsi2(in, &out.Kills)
		case "pilots":
			out.Pilots = int32(in.Int32())
		case "victory_points":
			easyjson9f888528DecodeGithubComAntihaxGoesiEsi3(in, &out.VictoryPoints)
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
func easyjson9f888528EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdFwStatsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"enlisted_on\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.EnlistedOn).MarshalJSON())
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if true {
		const prefix string = ",\"kills\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9f888528EncodeGithubComAntihaxGoesiEsi2(out, in.Kills)
	}
	if in.Pilots != 0 {
		const prefix string = ",\"pilots\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Pilots))
	}
	if true {
		const prefix string = ",\"victory_points\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9f888528EncodeGithubComAntihaxGoesiEsi3(out, in.VictoryPoints)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdFwStatsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f888528EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdFwStatsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f888528EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdFwStatsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f888528DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdFwStatsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f888528DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson9f888528DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetCorporationsCorporationIdFwStatsVictoryPoints) {
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
		case "last_week":
			out.LastWeek = int32(in.Int32())
		case "total":
			out.Total = int32(in.Int32())
		case "yesterday":
			out.Yesterday = int32(in.Int32())
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
func easyjson9f888528EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetCorporationsCorporationIdFwStatsVictoryPoints) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastWeek != 0 {
		const prefix string = ",\"last_week\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.LastWeek))
	}
	if in.Total != 0 {
		const prefix string = ",\"total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Total))
	}
	if in.Yesterday != 0 {
		const prefix string = ",\"yesterday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Yesterday))
	}
	out.RawByte('}')
}
func easyjson9f888528DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCorporationsCorporationIdFwStatsKills) {
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
		case "last_week":
			out.LastWeek = int32(in.Int32())
		case "total":
			out.Total = int32(in.Int32())
		case "yesterday":
			out.Yesterday = int32(in.Int32())
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
func easyjson9f888528EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCorporationsCorporationIdFwStatsKills) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastWeek != 0 {
		const prefix string = ",\"last_week\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.LastWeek))
	}
	if in.Total != 0 {
		const prefix string = ",\"total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Total))
	}
	if in.Yesterday != 0 {
		const prefix string = ",\"yesterday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Yesterday))
	}
	out.RawByte('}')
}
