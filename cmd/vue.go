package cmd

import (
	"fmt"
)

type Vue struct {
	Name       string
	Typescript bool
	Kind       VueKind
}

func NewVue(o Options) Vue {
	kind := Composition
	if o.Kind == "option" {
		kind = Option
	}
	return Vue{
		Name:       o.Name,
		Typescript: o.Typescript,
		Kind:       kind,
	}
}

func (v Vue) GetScript() string {
	ts := ` lang="ts"`
	if !v.Typescript {
		ts = ""
	}
	return fmt.Sprintf(`<script%s>
%s
</script>`, ts, v.GetCode())
}

func (v Vue) GetCode() string {
	if v.Typescript {
		return v.GetTs()
	} else {
		return v.GetJs()
	}
}

func (v Vue) GetJs() string {
	if v.Kind == Option {
		return `export default {};`
	} else {
		return `import {
  defineComponent,
} from '@vue/composition-api';

export default defineComponent({});`
	}
}

func (v Vue) GetTs() string {
	if v.Kind == Option {
		return `import Vue from 'vue';
export default Vue.extend({});`
	} else {
		return `import {
  defineComponent,
} from '@vue/composition-api';

export default defineComponent({});`
	}
}

func (v Vue) GetTemplate() string {
	return `<template>
  <div class="` + v.Name + `">
  </div>
</template>`
}

func (v Vue) GetStyle() string {
	return `<style lang="scss" scoped>
@use './` + v.Name + `.scss';
</style>`
}

func (v Vue) GetComponent() string {
	return fmt.Sprintf("%s\n\n%s\n\n%s", v.GetScript(), v.GetTemplate(), v.GetStyle())
}
