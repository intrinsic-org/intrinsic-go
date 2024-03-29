// This file was auto-generated by Fern from our API Definition.

package intrinsic

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/intrinsic-org/intrinsic-go/core"
)

type CreateEventTypeRequest struct {
	// Name of the event type to create.
	Name string `json:"name" url:"name"`
	// Fields of the event type
	Fields []*EventTypeField `json:"fields,omitempty" url:"fields,omitempty"`
}

type PatchEventTypeRequest struct {
	// Fields to add to the event type.
	Fields []*PatchEventTypeField `json:"fields,omitempty" url:"fields,omitempty"`
}

type EventTypeObject struct {
	// ID of the event type
	ID string `json:"id" url:"id"`
	// Name of the event type. Must be unique, and can only contain alphanumeric characters and underscores, and be up to 255 characters
	Name string `json:"name" url:"name"`
	// Fields of the event type
	Fields []*EventTypeField `json:"fields,omitempty" url:"fields,omitempty"`
	object string

	_rawJSON json.RawMessage
}

func (e *EventTypeObject) Object() string {
	return e.object
}

func (e *EventTypeObject) UnmarshalJSON(data []byte) error {
	type embed EventTypeObject
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EventTypeObject(unmarshaler.embed)
	e.object = "event_type"
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventTypeObject) MarshalJSON() ([]byte, error) {
	type embed EventTypeObject
	var marshaler = struct {
		embed
		Object string `json:"object"`
	}{
		embed:  embed(*e),
		Object: "event_type",
	}
	return json.Marshal(marshaler)
}

func (e *EventTypeObject) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ListEventTypesResponse struct {
	Data   []*EventTypeObject `json:"data,omitempty" url:"data,omitempty"`
	object string

	_rawJSON json.RawMessage
}

func (l *ListEventTypesResponse) Object() string {
	return l.object
}

func (l *ListEventTypesResponse) UnmarshalJSON(data []byte) error {
	type embed ListEventTypesResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListEventTypesResponse(unmarshaler.embed)
	l.object = "list"
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListEventTypesResponse) MarshalJSON() ([]byte, error) {
	type embed ListEventTypesResponse
	var marshaler = struct {
		embed
		Object string `json:"object"`
	}{
		embed:  embed(*l),
		Object: "list",
	}
	return json.Marshal(marshaler)
}

func (l *ListEventTypesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type PatchEventTypeField struct {
	// Name of the field. Must be unique, and can only contain alphanumeric characters and underscores, and be up to 255 characters
	FieldName string `json:"field_name" url:"field_name"`
	// Type of the field. Can be either strings, numbers, links to JPEG images, or links to video files
	Type EventTypeFieldType `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PatchEventTypeField) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchEventTypeField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchEventTypeField(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchEventTypeField) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
