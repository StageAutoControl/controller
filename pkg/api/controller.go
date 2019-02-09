package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

type Controller struct {
	logger         *logrus.Entry
	storage        storage
	controlledType interface{}
}

//type Arg json.RawMessage

type Response struct {
	value interface{}
}

type listResponse struct {
	value []interface{}
}

func newController(logger *logrus.Entry, storage storage, controlledType interface{}) *Controller {
	return &Controller{
		logger: logger,
		storage: storage,
		controlledType: controlledType,
	}
}

func (c *Controller) Create(r *http.Request, args *json.RawMessage, reply *Response) error {
	if args == nil || len(*args) == 0 {
		return errors.New("no parameter given")
	}

	t := reflect.ValueOf(c.controlledType).Type()
	target := reflect.New(t).Elem()
	if err := json.Unmarshal(*args, &target); err != nil {
		return fmt.Errorf("failed to unmarshal json content: %v", err)
	}

	id, err := c.getID(&target)
	if err != nil {
		return err
	}

	if err := c.storage.Write(id, &target); err != nil {
		return fmt.Errorf("failed to write to disk: %v", err)
	}

	reply.value = target
	return nil
}

func (c *Controller) checkType(value interface{}) error {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("parameter %s of type %s is no pointer", t.Name(), t.Kind())
	}

	s := t.Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("parameter pointer %v of type %v is no struct", s.Name(), s.Kind())
	}

	expected := reflect.TypeOf(c.controlledType)
	if s.Name() !=  expected.Name() {
		return fmt.Errorf("expected to get value of type %v, got type %v", expected.Name(), s.Name())
	}

	return nil
}

// getID expects a parameter "value" which is a pointer to a struct, which has a field called ID.
// if the ID field is set the value is returned, otherwise a new UUID v4 is generated and set as field value.
func (c *Controller) getID(value interface{}) (string, error) {
	v := reflect.ValueOf(value).Elem()
	field := v.FieldByName("ID")
	if !field.IsValid() {
		return "", fmt.Errorf("field ID of struct %s is not valid", v.Kind())
	}

	if !field.CanSet() {
		return "", fmt.Errorf("field ID of struct %s is not settable", v.Kind())
	}

	if field.Kind() != reflect.String {
		return "", fmt.Errorf("field ID of struct %s is not a string", v.Kind())
	}

	id := field.String()
	if id == "" {
		id = uuid.NewV4().String()
		field.SetString(id)
	}

	return id, nil
}
