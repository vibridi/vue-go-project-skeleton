## Vue+Go Project Skeleton
This is a project skeleton for a Vue.js+Typescript webapp on top of a Go backend. 
What you can do with it is to clone it and add or delete stuff based on the features you need in your project.   

### Technologies 

- [Vue.js v3](https://v3.vuejs.org/) - Core frontend framework
- [vue-class-component](https://class-component.vuejs.org/) - Library for Vue components in class style
- [vuex](https://vuex.vuejs.org/) - State management library for Vue (store pattern)
- [vuex-class](https://github.com/ktsn/vuex-class) - Binding helpers for Vuex and vue-class-component
- [vue-router](https://router.vuejs.org/) - Vue router for SPA development
- [Go 1.16](https://github.com/golang/go) - Well... Go
- [gin-gonic](https://github.com/gin-gonic/gin) - Web framework for Go

Webpack configuration lives in `vue.config.js`. See the [vue-cli](https://cli.vuejs.org/guide/) documentation 
for how to customize your webpack build.
  

### Getting started

Clone the project, then run:
```
$ yarn
```

Start the web server:
```
$ yarn dev
```

Update Go's dependencies:
```
$ make deps
```

### License
MIT