var app=function(){"use strict";function t(){}function n(t){return t()}function e(){return Object.create(null)}function o(t){t.forEach(n)}function s(t){return"function"==typeof t}function a(t,n){return t!=t?n==n:t!==n||t&&"object"==typeof t||"function"==typeof t}function r(t,n){t.appendChild(n)}function c(t,n,e){t.insertBefore(n,e||null)}function u(t){t.parentNode.removeChild(t)}function l(t){return document.createElement(t)}function i(){return t=" ",document.createTextNode(t);var t}function f(t,n,e,o){return t.addEventListener(n,e,o),()=>t.removeEventListener(n,e,o)}function p(t,n,e){null==e?t.removeAttribute(n):t.getAttribute(n)!==e&&t.setAttribute(n,e)}function d(t,n){t.value=null==n?"":n}let h;function m(t){h=t}const b=[],$=[],y=[],g=[],v=Promise.resolve();let x=!1;function _(t){y.push(t)}const w=new Set;let j=0;function k(){const t=h;do{for(;j<b.length;){const t=b[j];j++,m(t),z(t.$$)}for(m(null),b.length=0,j=0;$.length;)$.pop()();for(let t=0;t<y.length;t+=1){const n=y[t];w.has(n)||(w.add(n),n())}y.length=0}while(b.length);for(;g.length;)g.pop()();x=!1,w.clear(),m(t)}function z(t){if(null!==t.fragment){t.update(),o(t.before_update);const n=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,n),t.after_update.forEach(_)}}const O=new Set;function C(t,n){-1===t.$$.dirty[0]&&(b.push(t),x||(x=!0,v.then(k)),t.$$.dirty.fill(0)),t.$$.dirty[n/31|0]|=1<<n%31}function E(a,r,c,l,i,f,p,d=[-1]){const b=h;m(a);const $=a.$$={fragment:null,ctx:null,props:f,update:t,not_equal:i,bound:e(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(r.context||(b?b.$$.context:[])),callbacks:e(),dirty:d,skip_bound:!1,root:r.target||b.$$.root};p&&p($.root);let y=!1;if($.ctx=c?c(a,r.props||{},((t,n,...e)=>{const o=e.length?e[0]:n;return $.ctx&&i($.ctx[t],$.ctx[t]=o)&&(!$.skip_bound&&$.bound[t]&&$.bound[t](o),y&&C(a,t)),n})):[],$.update(),y=!0,o($.before_update),$.fragment=!!l&&l($.ctx),r.target){if(r.hydrate){const t=function(t){return Array.from(t.childNodes)}(r.target);$.fragment&&$.fragment.l(t),t.forEach(u)}else $.fragment&&$.fragment.c();r.intro&&((g=a.$$.fragment)&&g.i&&(O.delete(g),g.i(v))),function(t,e,a,r){const{fragment:c,on_mount:u,on_destroy:l,after_update:i}=t.$$;c&&c.m(e,a),r||_((()=>{const e=u.map(n).filter(s);l?l.push(...e):o(e),t.$$.on_mount=[]})),i.forEach(_)}(a,r.target,r.anchor,r.customElement),k()}var g,v;m(b)}function J(n){let e,s,a,h,m,b,$,y,g,v,x,_,w,j,k,z,O,C,E,J;return{c(){e=l("meta"),s=l("html"),a=i(),h=l("main"),m=l("div"),m.innerHTML='<h2 class="svelte-16ylzhb">Jsonata Transform Checker</h2>',b=i(),$=l("div"),y=l("textarea"),g=i(),v=l("textarea"),x=i(),_=l("textarea"),w=i(),j=l("div"),k=l("button"),k.textContent="Run Jsonata\n      ",z=l("button"),z.textContent="Escape Jsonata",O=i(),C=l("p"),C.textContent="Paste data in the input section, write some jsonata in middle bit & click a button",document.title="Jsonata Transform Checker",p(e,"name","robots"),p(e,"content","noindex nofollow"),p(e,"class","svelte-16ylzhb"),p(s,"lang","en"),p(s,"class","svelte-16ylzhb"),p(m,"class","header svelte-16ylzhb"),p(y,"class","column svelte-16ylzhb"),p(y,"rows","16"),p(y,"cols","50"),p(y,"placeholder","input JSON"),p(v,"class","column svelte-16ylzhb"),p(v,"rows","16"),p(v,"cols","50"),p(v,"placeholder","JSONATA"),p(_,"class","column svelte-16ylzhb"),p(_,"rows","16"),p(_,"cols","50"),p(_,"placeholder","output JSON"),p($,"class","row svelte-16ylzhb"),p(k,"class","svelte-16ylzhb"),p(z,"class","svelte-16ylzhb"),p(C,"class","svelte-16ylzhb"),p(j,"class","footer svelte-16ylzhb"),p(h,"class","svelte-16ylzhb")},m(t,o){r(document.head,e),r(document.head,s),c(t,a,o),c(t,h,o),r(h,m),r(h,b),r(h,$),r($,y),d(y,n[0].input),r($,g),r($,v),d(v,n[0].jsonata),r($,x),r($,_),d(_,n[0].output),r(h,w),r(h,j),r(j,k),r(j,z),r(j,O),r(j,C),E||(J=[f(y,"input",n[3]),f(v,"input",n[4]),f(_,"input",n[5]),f(k,"click",n[2]),f(z,"click",n[1])],E=!0)},p(t,[n]){1&n&&d(y,t[0].input),1&n&&d(v,t[0].jsonata),1&n&&d(_,t[0].output)},i:t,o:t,d(t){u(e),u(s),t&&u(a),t&&u(h),E=!1,o(J)}}}function A(t,n,e){let o={input:"",jsonata:"",output:""};return[o,function(){e(0,o.output=JSON.stringify(o.jsonata),o)},async function(){console.log("input: ",o.input);let t=await fetch("http://127.0.0.1:8050/jsonata",{method:"POST",headers:{"Content-Type":"application/json","Access-Control-Allow-Origin":"*"},body:JSON.stringify({input:o.input,jsonata:o.jsonata})}),n=await t.json();console.log("output: ",n),e(0,o.output=n.Output,o)},function(){o.input=this.value,e(0,o)},function(){o.jsonata=this.value,e(0,o)},function(){o.output=this.value,e(0,o)}]}return new class extends class{$destroy(){!function(t,n){const e=t.$$;null!==e.fragment&&(o(e.on_destroy),e.fragment&&e.fragment.d(n),e.on_destroy=e.fragment=null,e.ctx=[])}(this,1),this.$destroy=t}$on(t,n){const e=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return e.push(n),()=>{const t=e.indexOf(n);-1!==t&&e.splice(t,1)}}$set(t){var n;this.$$set&&(n=t,0!==Object.keys(n).length)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}{constructor(t){super(),E(this,t,A,J,a,{})}}({target:document.body,props:{name:"world"}})}();
//# sourceMappingURL=bundle.js.map
