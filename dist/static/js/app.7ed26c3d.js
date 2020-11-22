(function(t){function e(e){for(var n,s,i=e[0],l=e[1],c=e[2],u=0,d=[];u<i.length;u++)s=i[u],Object.prototype.hasOwnProperty.call(r,s)&&r[s]&&d.push(r[s][0]),r[s]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(t[n]=l[n]);f&&f(e);while(d.length)d.shift()();return o.push.apply(o,c||[]),a()}function a(){for(var t,e=0;e<o.length;e++){for(var a=o[e],n=!0,s=1;s<a.length;s++){var i=a[s];0!==r[i]&&(n=!1)}n&&(o.splice(e--,1),t=l(l.s=a[0]))}return t}var n={},s={app:0},r={app:0},o=[];function i(t){return l.p+"static/js/"+({}[t]||t)+"."+{"chunk-ca910568":"c40d5609"}[t]+".js"}function l(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,l),a.l=!0,a.exports}l.e=function(t){var e=[],a={"chunk-ca910568":1};s[t]?e.push(s[t]):0!==s[t]&&a[t]&&e.push(s[t]=new Promise((function(e,a){for(var n="static/css/"+({}[t]||t)+"."+{"chunk-ca910568":"7fd96764"}[t]+".css",r=l.p+n,o=document.getElementsByTagName("link"),i=0;i<o.length;i++){var c=o[i],u=c.getAttribute("data-href")||c.getAttribute("href");if("stylesheet"===c.rel&&(u===n||u===r))return e()}var d=document.getElementsByTagName("style");for(i=0;i<d.length;i++){c=d[i],u=c.getAttribute("data-href");if(u===n||u===r)return e()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=e,f.onerror=function(e){var n=e&&e.target&&e.target.src||r,o=new Error("Loading CSS chunk "+t+" failed.\n("+n+")");o.code="CSS_CHUNK_LOAD_FAILED",o.request=n,delete s[t],f.parentNode.removeChild(f),a(o)},f.href=r;var p=document.getElementsByTagName("head")[0];p.appendChild(f)})).then((function(){s[t]=0})));var n=r[t];if(0!==n)if(n)e.push(n[2]);else{var o=new Promise((function(e,a){n=r[t]=[e,a]}));e.push(n[2]=o);var c,u=document.createElement("script");u.charset="utf-8",u.timeout=120,l.nc&&u.setAttribute("nonce",l.nc),u.src=i(t);var d=new Error;c=function(e){u.onerror=u.onload=null,clearTimeout(f);var a=r[t];if(0!==a){if(a){var n=e&&("load"===e.type?"missing":e.type),s=e&&e.target&&e.target.src;d.message="Loading chunk "+t+" failed.\n("+n+": "+s+")",d.name="ChunkLoadError",d.type=n,d.request=s,a[1](d)}r[t]=void 0}};var f=setTimeout((function(){c({type:"timeout",target:u})}),12e4);u.onerror=u.onload=c,document.head.appendChild(u)}return Promise.all(e)},l.m=t,l.c=n,l.d=function(t,e,a){l.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},l.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},l.t=function(t,e){if(1&e&&(t=l(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(l.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)l.d(a,n,function(e){return t[e]}.bind(null,n));return a},l.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return l.d(e,"a",e),e},l.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},l.p="",l.oe=function(t){throw console.error(t),t};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],u=c.push.bind(c);c.push=e,c=c.slice();for(var d=0;d<c.length;d++)e(c[d]);var f=u;o.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},2395:function(t,e,a){},"4c04":function(t,e,a){"use strict";a.d(e,"b",(function(){return l})),a.d(e,"c",(function(){return c})),a.d(e,"d",(function(){return u})),a.d(e,"a",(function(){return d})),a.d(e,"e",(function(){return f}));a("d3b7"),a("96cf");var n=a("1da1"),s=(a("85e5"),a("bc3a")),r=a.n(s),o=a("a18c");a("5c96");r.a.defaults.baseURL="http://localhost:9090/",r.a.defaults.withCredentials=!0;var i=function(){var t=Object(n["a"])(regeneratorRuntime.mark((function t(e,a){var n,s,i,l,c=arguments;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:n=c.length>2&&void 0!==c[2]?c[2]:"GET",s=c.length>3?c[3]:void 0,n=n.toUpperCase(),i=!!s,{fullscreen:!0,background:"#010e2c8a",customClass:"loadingStyle"},r.a.interceptors.response.use((function(t){return 403!=t.data.code&&401!=t.data.code&&10005!=t.data.code||o["a"].push("/login"),t.data.success||t.data.message||(t.data.message="系统内部错误"),l&&l.close(),t}),(function(t){return Promise.reject(t)})),t.t0=n,t.next="GET"===t.t0?9:"POST"===t.t0?11:"POSTJSON"===t.t0?13:"POSTEXPORT"===t.t0?15:17;break;case 9:return t.abrupt("return",r.a.get(e,{params:a}));case 11:return t.abrupt("return",r.a.post(e,a,{headers:{"Content-Type":"application/x-www-form-urlencoded"}},i));case 13:return t.abrupt("return",r.a.post(e,a,{headers:{"Content-Type":"application/json"}}));case 15:return t.abrupt("return",r.a.post(e,a,{headers:{"Content-Type":"application/x-www-form-urlencoded"},responseType:"blob"}));case 17:return t.abrupt("break",18);case 18:case"end":return t.stop()}}),t)})));return function(e,a){return t.apply(this,arguments)}}(),l=(a("4328"),function(t){return i("/login",t,"POSTJSON")}),c=function(t){return i("/register",t,"POSTJSON")},u=function(t){return i("/tasklist",t,"GET")},d=function(t){return i("/issue",t,"POSTJSON")},f=function(t){return i("/update",t,"POSTJSON")}},"4ed6":function(t,e){t.exports="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAYAAAByDd+UAAABJUlEQVRIS83WMWvCQBTA8f+bHPwAQpdugkjp0tJNcOmHdhHcSicpIri5CH6ADp2evHLKmeaSdyZNzJhc7vfe5d3LCYCqDoAnYCsi33avrUtVh8AE+BKRHwnYB/AM7IF3Edm1AarqGFgAj8AaeDPwBfiMgAMwb4oGbAk8RHO/Gmgpb0IU52eN0ARmqzeV8A0t9WI0N6EJ7DLXL9gWWoeZcwGboh7sD3gr6sVKwVw0B0uCXjQXqwTr0FBr2ZV9VTRl3SWRxTGMHeU2jFqwItM4PveedYERugLirOyRZTvztsL7BBPf8X+WtNOiqdpnrW8Lz6b2jClutdKiyZkoZ2z/zTs32ni5vO/28wP2RlfWa4v36ubq/hDVxzHRTt3dHYTDn6Czo/4Jokgapo0Pe6wAAAAASUVORK5CYII="},5020:function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("e260"),a("e6cf"),a("cca6"),a("a79d");var n=a("2b0e"),s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("router-view")],1)},r=[],o=(a("7c55"),a("2877")),i={},l=Object(o["a"])(i,s,r,!1,null,null,null),c=l.exports,u=a("a18c"),d=a("2f62");n["default"].use(d["a"]);var f=new d["a"].Store({state:{menuActive:"1"},mutations:{},actions:{},modules:{}}),p=(a("9bf2a"),a("5c96")),m=a.n(p);a("0fae");n["default"].config.productionTip=!1,n["default"].use(m.a),new n["default"]({router:u["a"],store:f,render:function(t){return t(c)}}).$mount("#app")},"5fa3":function(t,e,a){"use strict";var n=a("8906"),s=a.n(n);s.a},"7c55":function(t,e,a){"use strict";var n=a("2395"),s=a.n(n);s.a},"85e5":function(t,e,a){"use strict";a.d(e,"b",(function(){return n})),a.d(e,"a",(function(){return s})),a.d(e,"c",(function(){return r}));a("baa5"),a("4d63"),a("ac1f"),a("25f0"),a("5319");var n=function(t,e){t&&("string"!==typeof e&&(e=JSON.stringify(e)),window.localStorage.setItem(t,e))},s=function(t){if(t)return window.localStorage.getItem(t)};function r(t){return 0===t?"未接收":1===t?"进行中":2===t?"已提交":3===t?"已确认":void 0}},8906:function(t,e,a){},"9bf2a":function(t,e){var a,n=document.querySelector("html");console.log(document.documentElement.clientWidth);var s=document.documentElement.clientWidth;a=100*s/1920,n.style.fontSize=a+"px",window.addEventListener("resize",(function(){var t=document.querySelector("html");console.log(document.documentElement.clientWidth);var e=document.documentElement.clientWidth,a=100*e/1920;t.style.fontSize=a+"px"}))},a18c:function(t,e,a){"use strict";a("d3b7");var n=a("2b0e"),s=a("8c4f"),r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"index"},[n("Header"),n("div",{staticClass:"con"},[n("div",{staticClass:"con_top"},[n("div",{staticClass:"con_top_btn",on:{click:t.openDialog}},[t._v(" 发布任务 ")])]),n("div",{staticClass:"con_box"},[n("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData,stripe:""}},[n("el-table-column",{attrs:{prop:"task_id",label:"任务ID"}}),n("el-table-column",{attrs:{prop:"issuer",label:"发布者"}}),n("el-table-column",{attrs:{prop:"task_user",label:"执行者"}}),n("el-table-column",{attrs:{prop:"bonus",label:"任务奖励"}}),n("el-table-column",{attrs:{prop:"task_name",label:"任务描述"}}),n("el-table-column",{attrs:{label:"进度"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t._f("taskStatus")(e.row.task_status))+" ")]}}])}),n("el-table-column",{attrs:{prop:"comment",label:"任务评价"}}),n("el-table-column",{attrs:{label:"动作"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("div",{staticClass:"check_btn",on:{click:function(a){return t.openDialog2(e.row)}}},[t._v("操作")])]}}])})],1)],1),n("div",{staticClass:"paging"},[n("el-pagination",{attrs:{"current-page":t.currentPage,"page-size":t.pageSize,layout:"total,  prev, pager, next, jumper",total:t.total},on:{"size-change":t.handleSizeChange,"current-change":t.handleCurrentChange}})],1)]),t.dialogFlag?n("div",{staticClass:"mask"},[n("div",{staticClass:"dialog"},[n("div",{staticClass:"dialog_top"},[n("div",{staticClass:"dialog_top_title"},[t._v("发布任务")]),n("div",{staticClass:"dialog_top_right"},[n("img",{attrs:{src:a("4ed6"),alt:""},on:{click:t.closeDialog}})])]),n("div",{staticClass:"dialog_con"},[n("el-form",{ref:"form",attrs:{model:t.taskData,"label-width":"1.3rem"}},[n("el-form-item",{attrs:{label:"任务描述"}},[n("el-input",{attrs:{placeholder:"请填写任务描述"},model:{value:t.taskData.task_name,callback:function(e){t.$set(t.taskData,"task_name",e)},expression:"taskData.task_name"}})],1),n("el-form-item",{attrs:{label:"任务奖励"}},[n("el-input",{attrs:{placeholder:"请填写任务奖励"},model:{value:t.taskData.bonus,callback:function(e){t.$set(t.taskData,"bonus",e)},expression:"taskData.bonus"}})],1),n("div",{staticClass:"sure_btn",on:{click:function(e){return t.issue(t.taskData)}}},[t._v("发布")])],1)],1)])]):t._e(),t.dialogFlag2?n("div",{staticClass:"mask"},[n("div",{staticClass:"dialog"},[n("div",{staticClass:"dialog_top"},[n("div",{staticClass:"dialog_top_title"},[t._v("操作")]),n("div",{staticClass:"dialog_top_right"},[n("img",{attrs:{src:a("4ed6"),alt:""},on:{click:t.closeDialog2}})])]),n("div",{staticClass:"dialog_con"},[n("el-form",{attrs:{model:t.statusData,"label-width":"1rem"}},[n("el-form-item",{attrs:{label:"状态"}},[n("el-select",{attrs:{placeholder:"请选择任务状态"},model:{value:t.statusData.task_status,callback:function(e){t.$set(t.statusData,"task_status",e)},expression:"statusData.task_status"}},[n("el-option",{staticStyle:{"font-size":"0.16rem"},attrs:{label:"接受",value:"1"}}),n("el-option",{staticStyle:{"font-size":"0.16rem"},attrs:{label:"提交",value:"2"}}),n("el-option",{staticStyle:{"font-size":"0.16rem"},attrs:{label:"确认",value:"3"}}),n("el-option",{staticStyle:{"font-size":"0.16rem"},attrs:{label:"退回",value:"4"}})],1)],1),n("el-form-item",{attrs:{label:"评价"}},[n("el-input",{attrs:{placeholder:"请填写任务评价"},model:{value:t.statusData.comment,callback:function(e){t.$set(t.statusData,"comment",e)},expression:"statusData.comment"}})],1),n("div",{staticClass:"sure_btn",on:{click:function(e){return t.update(t.statusData)}}},[t._v("确定")])],1)],1)])]):t._e()],1)},o=[],i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"head"},[a("div",{staticClass:"head_left"},[t._v(" tokentask ")]),a("div",{staticClass:"head_right"},[t._v(" hi，"+t._s(t.username)+" ")])])},l=[],c=a("85e5"),u={data:function(){return{username:""}},created:function(){this.username=Object(c["a"])("username")}},d=u,f=(a("d0ed"),a("2877")),p=Object(f["a"])(d,i,l,!1,null,"82b34970",null),m=p.exports,g=a("4c04"),h={created:function(){this.initData()},filters:{taskStatus:function(t){return Object(c["c"])(t)}},data:function(){return{lectureName:"",currentPage:1,pageSize:10,total:0,dialogFlag:!1,dialogFlag2:!1,tableData:[],taskData:{task_name:"",task_bonus:""},statusData:{task_id:"",task_status:"",comment:""}}},methods:{initData:function(){var t=this;Object(g["d"])({page:this.currentPage}).then((function(e){"0"==e.data.code?(t.tableData=e.data.data.data,t.total=e.data.data.total):t.$message.error(e.data.msg)}))},issue:function(t){var e=this;if(/^\s*$/g.test(t.task_name))this.$message({message:"请填写任务描述",type:"error"});else if(/^\s*$/g.test(t.bonus))this.$message.error("请填写任务奖励");else{var a={task_name:t.task_name,bonus:parseInt(t.bonus)};console.log(a),Object(g["a"])(a).then((function(t){console.log(t),"0"==t.data.code?(e.$message.success("发布成功"),e.dialogFlag=!1,e.initData(),e.taskData.task_name="",e.taskData.bonus=""):e.$message.error(t.data.msg)}))}},update:function(t){var e,a=this;t.task_status||this.$message.error("请选择任务状态"),e=t.comment?{task_id:t.task_id,task_status:parseInt(t.task_status),comment:t.comment}:{task_id:t.task_id,task_status:parseInt(t.task_status)},Object(g["e"])(e).then((function(t){"0"==t.data.code?(a.initData(),a.$message.success("操作成功"),a.dialogFlag2=!1,a.statusData.task_status="",a.statusData.comment=""):a.$message.error(t.data.msg)}))},openDialog:function(){this.dialogFlag=!0},closeDialog:function(){this.dialogFlag=!1},openDialog2:function(t){this.dialogFlag2=!0,this.statusData.task_id=t.task_id},closeDialog2:function(){this.dialogFlag2=!1},handleSizeChange:function(t){},handleCurrentChange:function(t){this.currentPage=t,this.initData()}},components:{Header:m}},b=h,v=(a("5fa3"),Object(f["a"])(b,r,o,!1,null,"0a77642e",null)),_=v.exports;n["default"].use(s["a"]);var k=[{path:"/login",name:"Login",component:function(){return a.e("chunk-ca910568").then(a.bind(null,"a55b"))},meta:{title:"登录"}},{path:"/",component:_,redirect:"/index",children:[{path:"index",name:"Index",component:_,meta:{requireAuth:!0,title:"首页"}}]}],y=new s["a"]({mode:"history",base:"",routes:k});e["a"]=y;y.beforeEach((function(t,e,a){t.meta.title&&(document.title=t.meta.title),t.meta.requireAuth?Object(c["a"])("username")?a():a({path:"/login"}):a()}))},d0ed:function(t,e,a){"use strict";var n=a("5020"),s=a.n(n);s.a}});