package routeros

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// CoverageField returns the schema field a RouterOS property is read into,
// applying the resource's transform set and the attribute drift for the given
// RouterOS version, and whether the resource skips the property.
// Used by tools/ros-coverage.
func CoverageField(s map[string]*schema.Schema, ros, property string) (field string, skipped bool) {
	name := property
	if ts, ok := s[MetaTransformSet]; ok {
		if v, ok := loadTransformSet(ts.Default.(string), true)[name]; ok {
			name = v
		}
	}
	if v, ok := driftAttributeSlice.GetDriftMap(ros, s[MetaResourcePath].Default.(string), true)[name]; ok {
		name = v
	}
	field = KebabToSnake(name)
	if sf, ok := s[MetaSkipFields]; ok {
		_, skipped = loadSkipFields(sf.Default.(string))[field]
	}
	return
}

// CoverageProperty returns the RouterOS property a schema field is written to.
// Used by tools/ros-coverage.
func CoverageProperty(s map[string]*schema.Schema, ros, field string) string {
	if ts, ok := s[MetaTransformSet]; ok {
		if v, ok := loadTransformSet(ts.Default.(string), false)[field]; ok {
			return v
		}
	}
	if v, ok := driftAttributeSlice.GetDriftMap(ros, s[MetaResourcePath].Default.(string), false)[field]; ok {
		return v
	}
	return SnakeToKebab(field)
}
