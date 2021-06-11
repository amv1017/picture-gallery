(this["webpackJsonpart-gallery-view"]=this["webpackJsonpart-gallery-view"]||[]).push([[0],{24:function(e,t,n){},32:function(e,t,n){"use strict";n.r(t);var r=n(1),c=n.n(r),a=n(18),s=n.n(a),i=(n(24),n(4)),l=n.n(i),o=n(9),u=n(5),d=n(0),j=function(e){var t=e.paint,n=e.onClick;return Object(d.jsx)("div",{className:"max-w-sm rounded overflow-hidden shadow-lg bg-gray-200",onClick:n,children:Object(d.jsxs)("div",{className:"px-6 py-4",children:[Object(d.jsx)("img",{className:"w-full",src:t.url}),Object(d.jsx)("p",{className:"text-center",children:t.title})]})})},x=function(e){var t=e.id,n=Object(r.useState)({}),c=Object(u.a)(n,2),a=c[0],s=c[1];Object(r.useEffect)((function(){(function(){var e=Object(o.a)(l.a.mark((function e(){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.t0=s,e.next=3,i();case 3:e.t1=e.sent,(0,e.t0)(e.t1);case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}})()()}),[]);var i=function(){var e=Object(o.a)(l.a.mark((function e(){var n,r;return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://localhost:8080/api/painting/".concat(t),{method:"GET"});case 2:return n=e.sent,e.next=5,n.json();case 5:return r=e.sent,e.abrupt("return",r);case 7:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),j=function(){var e=Object(o.a)(l.a.mark((function e(){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://localhost:8080/api/painting/".concat(t),{method:"DELETE"});case 2:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return Object(d.jsx)("div",{children:Object(d.jsxs)("div",{className:"flex flex-col justify-center",children:[Object(d.jsx)("img",{className:"object-contain h-screen w-full",src:a.url}),Object(d.jsx)("p",{className:"text-blue-900 text-4xl text-center mb-4",children:a.title}),Object(d.jsx)("p",{className:"text-blue-900",children:a.description}),Object(d.jsx)("div",{className:"flex justify-end",children:Object(d.jsx)("button",{className:"bg-red-400 rounded-md",onClick:function(){return j()},children:"Delete This Exhibit"})})]})})},b=function(e){var t=e.gallery,n=Object(r.useState)(!1),c=Object(u.a)(n,2),a=c[0],s=c[1],i=Object(r.useState)(1),l=Object(u.a)(i,2),o=l[0],b=l[1];return Object(d.jsx)(d.Fragment,{children:null===t?"No Exhibits":a?Object(d.jsx)("div",{className:"grid grid-cols-1 gap-4",children:Object(d.jsx)("div",{onClick:function(){return s(!a)},children:Object(d.jsx)(x,{id:t[o].id})})}):Object(d.jsx)("div",{className:"grid grid-cols-3 gap-4",children:t.map((function(e){return Object(d.jsx)("div",{onClick:function(){s(!a)},children:Object(d.jsx)(j,{paint:e,onClick:function(){b(t.indexOf(e))}})})}))})})},m=function(e){var t=Object(r.useState)([]),n=Object(u.a)(t,2),c=n[0],a=n[1];return Object(r.useEffect)((function(){(function(){var t=Object(o.a)(l.a.mark((function t(){var n,r;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,fetch("http://localhost:8080/api/".concat(e.match.params.section,"/").concat(e.match.params.id),{method:"GET"});case 2:return n=t.sent,t.next=5,n.json();case 5:r=t.sent,a(r);case 7:case"end":return t.stop()}}),t)})));return function(){return t.apply(this,arguments)}})()()}),[e.match.params.id]),Object(d.jsx)("div",{class:"w-full max-w-6xl mx-auto px-4 border-2 border-grey-600",children:Object(d.jsx)(b,{gallery:c})})},h=n(10),p=n(2),f=function(){var e=Object(r.useState)([]),t=Object(u.a)(e,2),n=t[0],c=t[1];Object(r.useEffect)((function(){(function(){var e=Object(o.a)(l.a.mark((function e(){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.t0=c,e.next=3,a();case 3:e.t1=e.sent,(0,e.t0)(e.t1);case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}})()()}),[]);var a=function(){var e=Object(o.a)(l.a.mark((function e(){var t,n;return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://localhost:8080/api/authors",{method:"GET"});case 2:return t=e.sent,e.next=5,t.json();case 5:return n=e.sent,e.abrupt("return",n);case 7:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return Object(d.jsxs)("div",{children:[Object(d.jsxs)("div",{className:"mb-3 text-center bg-blue-200 rounded-md p-3",children:["AUTHORS",Object(d.jsx)("div",{className:"flex flex-col items-center pb-6",children:n.map((function(e){return Object(d.jsx)("div",{children:Object(d.jsx)(h.b,{className:"bg-blue-400 rounded-md px-5 mb-4",to:"/author/"+e.id,children:e.name})})}))})]}),Object(d.jsxs)("div",{className:"mb-3 text-center bg-green-200 rounded-md p-3",children:["GENRES",Object(d.jsxs)("div",{className:"flex flex-col items-center",children:[Object(d.jsx)(h.b,{className:"bg-green-400 rounded-md py-2 px-5 mb-1",to:"/genre/1",children:"Landscape"}),Object(d.jsx)(h.b,{className:"bg-green-400 rounded-md py-2 px-5 mb-1",to:"/genre/2",children:"Portrait"}),Object(d.jsx)(h.b,{className:"bg-green-400 rounded-md py-2 px-5",to:"/genre/3",children:"Still Life"})]})]})]})};var O=function(){return Object(d.jsxs)("div",{className:"flex items-center h-8 bg-gradient-to-r from-purple-400 via-cyan-500 to-blue-500",children:[Object(d.jsx)("div",{className:"ml-16 text-black",children:"ART GALLERY"}),Object(d.jsxs)("ul",{className:"flex flex-auto justify-end pr-40",children:[Object(d.jsx)("li",{className:"text-gray-100 m-9",children:Object(d.jsx)(h.b,{to:"/",children:"Home"})}),Object(d.jsx)("li",{className:"text-gray-100 m-9",children:Object(d.jsx)(h.b,{to:"/new",children:"New"})})]})]})},v=function(){var e=Object(r.useState)(""),t=Object(u.a)(e,2),n=t[0],c=t[1],a=Object(r.useState)(""),s=Object(u.a)(a,2),i=s[0],j=s[1],x=Object(r.useState)(""),b=Object(u.a)(x,2),m=b[0],h=b[1],p=Object(r.useState)(""),f=Object(u.a)(p,2),O=f[0],v=f[1],g=Object(r.useState)(""),y=Object(u.a)(g,2),N=y[0],w=y[1],E=function(){var e=Object(o.a)(l.a.mark((function e(){var t;return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=JSON.stringify({title:n,description:i,url:m,genre:O,author:N}),e.next=3,fetch("http://localhost:8080/api/paintings",{method:"POST",headers:{"Content-Type":"application/json"},body:t});case 3:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return Object(d.jsxs)("div",{className:"flex flex-col",children:[Object(d.jsxs)("form",{className:"ml-40 mt-16 flex flex-col items-right",onSubmit:function(e){e.preventDefault(),E(),c(""),j(""),h(""),v(""),w("")},children:[Object(d.jsxs)("div",{children:[Object(d.jsx)("label",{htmlFor:"title",children:"Title"}),Object(d.jsx)("input",{type:"text",className:"w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4",value:n,placeholder:"Enter Title",onChange:function(e){return c(e.target.value)}})]}),Object(d.jsxs)("div",{children:[Object(d.jsx)("label",{htmlFor:"author",children:"Author"}),Object(d.jsx)("input",{type:"text",className:"w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4",value:N,placeholder:"Enter Author",onChange:function(e){return w(e.target.value)}})]}),Object(d.jsxs)("div",{children:[Object(d.jsx)("label",{htmlFor:"genre",children:"Genre"}),Object(d.jsx)("input",{type:"text",className:"w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4",value:O,placeholder:"Enter Genre",onChange:function(e){return v(e.target.value)}})]}),Object(d.jsxs)("div",{children:[Object(d.jsx)("label",{htmlFor:"description",children:"Description"}),Object(d.jsx)("input",{type:"text",className:"w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4",value:i,placeholder:"Enter Description",onChange:function(e){return j(e.target.value)}})]}),Object(d.jsxs)("div",{children:[Object(d.jsx)("label",{htmlFor:"url",children:"URL"}),Object(d.jsx)("input",{type:"text",className:"w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4",value:m,placeholder:"Enter URL",onChange:function(e){return h(e.target.value)}})]}),Object(d.jsx)("div",{className:"flex justify-center items-center mt-6",children:Object(d.jsx)("button",{className:"bg-green-300 py-2 px-4 text-sm text-white rounded border border-green focus:outline-none focus:border-green-dark",children:"Confirm"})})]}),Object(d.jsx)("div",{className:"ml-40 mt-16 flex flex-col items-right",children:Object(d.jsxs)("ul",{children:["Example",Object(d.jsx)("li",{children:"Title: Black Square"}),Object(d.jsx)("li",{children:"Author: Kazimir Malevich"}),Object(d.jsx)("li",{children:"Genre: Still Life"}),Object(d.jsx)("li",{children:"Description: Masterpiece..."}),Object(d.jsx)("li",{children:"URL: https://img.gazeta.ru/files3/907/7888907/malevich21_-pic905-895x505-1560.jpg"})]})})]})},g=function(){return Object(d.jsxs)("div",{className:"p-80 w-9/12 bg-gradient-to-r from-gray-400 via-gray-500 to-gray-800 rounded-md",children:[Object(d.jsx)("div",{className:"max-w-md text-center text-6xl text-gray-300",children:"WELCOME"}),Object(d.jsx)("div",{className:"max-w-md text-center text-2xl text-gray-400",children:"Impressive Art Gallery"}),Object(d.jsx)("div",{})]})};var y=function(){return Object(d.jsx)("div",{className:"container mx-auto",children:Object(d.jsxs)(h.a,{children:[Object(d.jsx)(O,{}),Object(d.jsxs)("div",{className:"flex flex-col md:flex-row",children:[Object(d.jsx)(f,{}),Object(d.jsxs)(p.c,{children:[Object(d.jsx)(p.a,{exact:!0,path:"/new",component:v}),Object(d.jsx)(p.a,{path:"/:section/:id",component:m}),Object(d.jsx)(p.a,{exact:!0,path:"/",component:g})]})]})]})})};s.a.render(Object(d.jsx)(c.a.StrictMode,{children:Object(d.jsx)(y,{})}),document.getElementById("root"))}},[[32,1,2]]]);
//# sourceMappingURL=main.6cf321fe.chunk.js.map