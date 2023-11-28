import{S as A,i as Q,s as R,y as g,a as S,k as E,z as d,c as q,l as I,m as B,h as v,n as C,A as y,b as N,D as h,N as V,Q as z,E as G,g as _,d as b,B as k}from"../chunks/index.f49fd5c2.js";import{F as H,a as U}from"../chunks/FormInputTextArea.2980ceeb.js";import{F as P}from"../chunks/FormInputField.6e32f34e.js";import{P as j,B as J,g as T}from"../chunks/PageHeader.b42d5f8a.js";import{k as K,m as W}from"../chunks/svelte-apollo.e94973ad.js";const X=K`
  mutation CreateLookout(
    $name: String!
    $query: String!
    $cron: String!
    $notifyLocal: Boolean!
    $notifyMail: Boolean!
  ) {
    createLookout(
      data: {
        name: $name
        query: $query
        cron: $cron
        notifyMail: $notifyMail
        notifyLocal: $notifyLocal
      }
    ) {
      id
    }
  }
`,Y=()=>W(X);function Z(x){let o,w,a,e,n,f,r,c,l,L,u,$,m,i,M,F,O;return o=new j({props:{title:"Create Lookout",backAction:x[1]}}),n=new P({props:{name:"name",required:!0,label:"Lookout name",placeholder:"My important query"}}),r=new P({props:{name:"cron",required:!0,label:"Cron expression",placeholder:"0 7 * * 1-5"}}),l=new H({props:{name:"query",required:!0,label:"SQL Query",placeholder:"SELECT COUNT(*) FROM something"}}),u=new U({props:{label:"Notifications",items:[{name:"notifyLocal",title:"Local notification",description:"Only works if run locally"},{name:"notifyMail",title:"Mail notification"}]}}),i=new J({props:{title:"Create",type:"submit",leadingIcon:"check"}}),{c(){g(o.$$.fragment),w=S(),a=E("form"),e=E("div"),g(n.$$.fragment),f=S(),g(r.$$.fragment),c=S(),g(l.$$.fragment),L=S(),g(u.$$.fragment),$=S(),m=E("div"),g(i.$$.fragment),this.h()},l(t){d(o.$$.fragment,t),w=q(t),a=I(t,"FORM",{method:!0,class:!0});var p=B(a);e=I(p,"DIV",{class:!0});var s=B(e);d(n.$$.fragment,s),f=q(s),d(r.$$.fragment,s),c=q(s),d(l.$$.fragment,s),L=q(s),d(u.$$.fragment,s),$=q(s),m=I(s,"DIV",{class:!0});var D=B(m);d(i.$$.fragment,D),D.forEach(v),s.forEach(v),p.forEach(v),this.h()},h(){C(m,"class","mt-4 mx-3 w-full flex flex-row-reverse"),C(e,"class","flex flex-wrap -mx-3 mb-6"),C(a,"method","POST"),C(a,"class","w-full max-w-lg")},m(t,p){y(o,t,p),N(t,w,p),N(t,a,p),h(a,e),y(n,e,null),h(e,f),y(r,e,null),h(e,c),y(l,e,null),h(e,L),y(u,e,null),h(e,$),h(e,m),y(i,m,null),M=!0,F||(O=V(a,"submit",z(x[0])),F=!0)},p:G,i(t){M||(_(o.$$.fragment,t),_(n.$$.fragment,t),_(r.$$.fragment,t),_(l.$$.fragment,t),_(u.$$.fragment,t),_(i.$$.fragment,t),M=!0)},o(t){b(o.$$.fragment,t),b(n.$$.fragment,t),b(r.$$.fragment,t),b(l.$$.fragment,t),b(u.$$.fragment,t),b(i.$$.fragment,t),M=!1},d(t){k(o,t),t&&v(w),t&&v(a),k(n),k(r),k(l),k(u),k(i),F=!1,O()}}}function tt(x){const o=Y();return[async e=>{var $,m,i;const n=new FormData(e.target),f=($=n.get("name"))==null?void 0:$.toString(),r=(m=n.get("cron"))==null?void 0:m.toString(),c=(i=n.get("query"))==null?void 0:i.toString(),l=n.has("notifyLocal"),L=n.has("notifyMail");if(!f||!r||!c){alert("Name, cron and query must not be null");return}if((await o({variables:{name:f,cron:r,query:c,notifyMail:L,notifyLocal:l}})).errors){alert("Something went wrong");return}T("/lookout",{state:{refetch:!0}})},()=>T("/lookout")]}class it extends A{constructor(o){super(),Q(this,o,tt,Z,R,{})}}export{it as default};
