package monday

import (
	"fmt"
	"log"
	"strings"
)

type Groups struct {
	fields []GroupsField
	args   []GroupsArgument
}

func (g Groups) stringify() string {
	fields := make([]string, 0)
	for _, field := range g.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range g.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`groups{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`groups(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))

}

func NewGroups(fields []GroupsField) Groups {
	if len(fields) == 0 {
		return Groups{
			fields: []GroupsField{
				GroupsIDField(),
			},
		}
	}

	return Groups{
		fields: fields,
	}
}

func NewGroupWithArguments(fields []GroupsField, args []GroupsArgument) Groups {
	groups := NewGroups(fields)
	groups.args = args
	return groups
}

type GroupsField struct {
	field string
	value interface{}
}

var (
	groupsArchivedField = GroupsField{"archived", nil}
	groupsColorField    = GroupsField{"color", nil}
	groupsDeletedField  = GroupsField{"deleted", nil}
	groupsIDField       = GroupsField{"id", nil}
	groupsPositionField = GroupsField{"position", nil}
	groupsTitleField    = GroupsField{"title", nil}
)

func (f GroupsField) stringify() string {
	switch f.field {
	case "items":
		return f.value.(Items).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

func GroupsArchivedField() GroupsField {
	return groupsArchivedField
}

func GroupsColorField() GroupsField {
	return groupsColorField
}

func GroupsDeletedField() GroupsField {
	return groupsDeletedField
}

func GroupsIDField() GroupsField {
	return groupsIDField
}

func NewGroupsItemsField(items Items) GroupsField {
	return GroupsField{field: "items", value: items}
}

func GroupsPositionField() GroupsField {
	return groupsPositionField
}

func GroupsTitleField() GroupsField {
	return groupsTitleField
}

type GroupsArgument struct {
	argument string
	value    interface{}
}

func (a GroupsArgument) stringify() string {
	switch a.argument {
	case "ids":
		switch ids := a.value.([]int); {
		case len(ids) == 1:
			return fmt.Sprintf("ids:%d", ids[0])
		case len(ids) > 1:
			return fmt.Sprintf("ids:%s", strings.Replace(fmt.Sprint(ids), " ", ",", -1))
		default:
			return ""
		}
	default:
		log.Fatalln("unreachable boards argument")
	}
	return ""
}

func NewIDsGroupsArg(ids []int) GroupsArgument {
	return GroupsArgument{
		argument: "ids",
		value:    ids,
	}
}
