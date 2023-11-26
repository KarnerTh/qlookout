import{S as oe,i as ue,s as se,y as I,a as x,e as Y,z as k,c as B,A as N,b as G,d as $,f as ie,g as h,B as D,h as R,F as Z,v as me,w as ee,R as te,k as A,l as L,m as O,n as C,D as y,N as fe,Q as pe,T as ne,E as de}from"../chunks/index.f49fd5c2.js";import{P as ce,g as W,B as ge}from"../chunks/PageHeader.8297dd43.js";import{p as $e}from"../chunks/stores.aa56ac9d.js";import{F as le,R as he}from"../chunks/RuleInputs.63f4425d.js";import{F as re}from"../chunks/FormInputField.6e32f34e.js";import{k as X,q as _e,m as ae}from"../chunks/svelte-apollo.d2e69eac.js";import{L as we}from"../chunks/LoadingSpinner.1a458fe5.js";import{D as be}from"../chunks/DialogButton.2fc8da89.js";const Te=X`
  query Rule($id: Int!) {
    rule(id: $id) {
      id
      lookoutId
      columnName
      columnType
      rowIndex
      exactValue
      greaterThan
      lessThan
      shouldBeNull
    }
  }
`,ye=u=>_e(Te,{variables:{id:u}}),ve=X`
  mutation UpdateRule(
    $id: Int!
    $columnName: String
    $columnType: String
    $rowIndex: Int!
    $exactValue: String
    $lessThan: String
    $greaterThan: String
    $shouldBeNull: Boolean
  ) {
    updateRule(
      id: $id
      data: {
        columnName: $columnName
        columnType: $columnType
        rowIndex: $rowIndex
        exactValue: $exactValue
        lessThan: $lessThan
        greaterThan: $greaterThan
        shouldBeNull: $shouldBeNull
      }
    ) {
      id
    }
  }
`,Ie=()=>ae(ve),ke=X`
  mutation DeleteRule(
    $id: Int!
  ) {
    deleteRule(
      id: $id
    ) {
      id
    }
  }
`,Ne=()=>ae(ke);function De(u){let o,e,t,l,a,i,m,c,S,n,f,_,w,q,V,E,b,v,g,p,F,M,Q;t=new re({props:{name:"columnName",value:u[2].data.rule.columnName,required:!0,label:"Column name",placeholder:"columnName"}}),a=new re({props:{name:"rowIndex",value:u[2].data.rule.rowIndex,type:"number",required:!0,label:"Row index",placeholder:"0"}});function J(r){u[9](r)}let j={name:"columnType",required:!0,label:"Column Type",options:[{name:"text",title:"text"},{name:"int",title:"int"},{name:"float",title:"float"}]};u[1]!==void 0&&(j.value=u[1]),m=new le({props:j}),ee.push(()=>te(m,"value",J));function K(r){u[10](r)}let z={name:"ruleType",required:!0,label:"Rule Type",options:[{name:"exact",title:"exact value"},{name:"null",title:"should be null"},{name:"less",title:"less than"},{name:"greater",title:"greater than"},{name:"between",title:"between"}]};return u[0]!==void 0&&(z.value=u[0]),n=new le({props:z}),ee.push(()=>te(n,"value",K)),w=new he({props:{ruleType:u[0],value:u[2].data.rule,inputType:u[3]}}),v=new be({props:{buttonText:"Delete",buttonIcon:"trash",buttonType:"red",dialogTitle:"Delete Rule",dialogDescription:"Do you realy want to delete this rule? This action cannot be undone",actionButtons:[{title:"Cancel",onClick:xe},{title:"Delete",buttonType:"red",onClick:u[11]}]}}),p=new ge({props:{title:"Update",type:"submit",leadingIcon:"check"}}),{c(){o=A("form"),e=A("div"),I(t.$$.fragment),l=x(),I(a.$$.fragment),i=x(),I(m.$$.fragment),S=x(),I(n.$$.fragment),_=x(),I(w.$$.fragment),q=x(),V=A("div"),E=x(),b=A("div"),I(v.$$.fragment),g=x(),I(p.$$.fragment),this.h()},l(r){o=L(r,"FORM",{method:!0,class:!0});var d=O(o);e=L(d,"DIV",{class:!0});var s=O(e);k(t.$$.fragment,s),l=B(s),k(a.$$.fragment,s),i=B(s),k(m.$$.fragment,s),S=B(s),k(n.$$.fragment,s),_=B(s),k(w.$$.fragment,s),q=B(s),V=L(s,"DIV",{class:!0}),O(V).forEach(R),E=B(s),b=L(s,"DIV",{class:!0});var T=O(b);k(v.$$.fragment,T),g=B(T),k(p.$$.fragment,T),T.forEach(R),s.forEach(R),d.forEach(R),this.h()},h(){C(V,"class","w-full"),C(b,"class","mt-4 mx-3 w-full flex flex-row-reverse gap-2"),C(e,"class","flex flex-wrap -mx-3 mb-6"),C(o,"method","POST"),C(o,"class","w-full max-w-lg")},m(r,d){G(r,o,d),y(o,e),N(t,e,null),y(e,l),N(a,e,null),y(e,i),N(m,e,null),y(e,S),N(n,e,null),y(e,_),N(w,e,null),y(e,q),y(e,V),y(e,E),y(e,b),N(v,b,null),y(b,g),N(p,b,null),F=!0,M||(Q=fe(o,"submit",pe(u[6])),M=!0)},p(r,d){const s={};d&4&&(s.value=r[2].data.rule.columnName),t.$set(s);const T={};d&4&&(T.value=r[2].data.rule.rowIndex),a.$set(T);const U={};!c&&d&2&&(c=!0,U.value=r[1],ne(()=>c=!1)),m.$set(U);const P={};!f&&d&1&&(f=!0,P.value=r[0],ne(()=>f=!1)),n.$set(P);const H={};d&1&&(H.ruleType=r[0]),d&4&&(H.value=r[2].data.rule),d&8&&(H.inputType=r[3]),w.$set(H)},i(r){F||(h(t.$$.fragment,r),h(a.$$.fragment,r),h(m.$$.fragment,r),h(n.$$.fragment,r),h(w.$$.fragment,r),h(v.$$.fragment,r),h(p.$$.fragment,r),F=!0)},o(r){$(t.$$.fragment,r),$(a.$$.fragment,r),$(m.$$.fragment,r),$(n.$$.fragment,r),$(w.$$.fragment,r),$(v.$$.fragment,r),$(p.$$.fragment,r),F=!1},d(r){r&&R(o),D(t),D(a),D(m),D(n),D(w),D(v),D(p),M=!1,Q()}}}function Re(u){let o,e,t;return e=new we({}),{c(){o=A("div"),I(e.$$.fragment),this.h()},l(l){o=L(l,"DIV",{class:!0});var a=O(o);k(e.$$.fragment,a),a.forEach(R),this.h()},h(){C(o,"class","flex justify-center pt-9")},m(l,a){G(l,o,a),N(e,o,null),t=!0},p:de,i(l){t||(h(e.$$.fragment,l),t=!0)},o(l){$(e.$$.fragment,l),t=!1},d(l){l&&R(o),D(e)}}}function Se(u){let o,e,t,l,a,i;o=new ce({props:{title:"Update Rule",backAction:u[8]}});const m=[Re,De],c=[];function S(n,f){return n[2].loading?0:n[2].data?1:-1}return~(t=S(u))&&(l=c[t]=m[t](u)),{c(){I(o.$$.fragment),e=x(),l&&l.c(),a=Y()},l(n){k(o.$$.fragment,n),e=B(n),l&&l.l(n),a=Y()},m(n,f){N(o,n,f),G(n,e,f),~t&&c[t].m(n,f),G(n,a,f),i=!0},p(n,[f]){let _=t;t=S(n),t===_?~t&&c[t].p(n,f):(l&&(me(),$(c[_],1,1,()=>{c[_]=null}),ie()),~t?(l=c[t],l?l.p(n,f):(l=c[t]=m[t](n),l.c()),h(l,1),l.m(a.parentNode,a)):l=null)},i(n){i||(h(o.$$.fragment,n),h(l),i=!0)},o(n){$(o.$$.fragment,n),$(l),i=!1},d(n){D(o,n),n&&R(e),~t&&c[t].d(n),n&&R(a)}}}const xe=function(){};function Be(u,o,e){let t,l;Z(u,$e,g=>e(12,l=g));let a,i,m;const c=Ie(),S=Ne(),n=+l.params.lookoutId,f=+l.params.ruleId,_=ye(f);Z(u,_,g=>e(2,t=g));const w=async g=>{var d,s,T,U,P;const p=new FormData(g.target),F=(d=p.get("columnName"))==null?void 0:d.toString(),M=(s=p.get("columnType"))==null?void 0:s.toString(),Q=+(p.get("rowIndex")??0),J=(T=p.get("exactValue"))==null?void 0:T.toString(),j=(U=p.get("lessThan"))==null?void 0:U.toString(),K=(P=p.get("greaterThan"))==null?void 0:P.toString(),z=p.has("shouldBeNull");if((await c({variables:{id:f,columnName:F,columnType:M,rowIndex:Q,exactValue:J,lessThan:j,greaterThan:K,shouldBeNull:z}})).errors){alert("Something went wrong");return}W(`/lookout/${n}`,{state:{refetch:!0}})},q=async()=>{if((await S({variables:{id:f}})).errors){alert("Something went wrong");return}W(`/lookout/${n}`,{state:{refetch:!0}})},V=()=>W(`/lookout/${n}`);function E(g){i=g,e(1,i),e(2,t),e(0,a)}function b(g){a=g,e(0,a),e(1,i),e(2,t)}const v=function(){q()};return u.$$.update=()=>{if(u.$$.dirty&7&&(i===void 0&&t.data&&e(1,i=t.data.rule.columnType),a===void 0&&t.data&&(t.data.rule.exactValue!==null?e(0,a="exact"):t.data.rule.shouldBeNull===!0?e(0,a="null"):t.data.rule.lessThan!==null&&t.data.rule.greaterThan!==null?e(0,a="between"):t.data.rule.lessThan!==null?e(0,a="less"):t.data.rule.greaterThan!==null&&e(0,a="greater"))),u.$$.dirty&2)e:{if(!i)break e;["int","float"].includes(i)?e(3,m="number"):e(3,m="text")}},[a,i,t,m,n,_,w,q,V,E,b,v]}class Ae extends oe{constructor(o){super(),ue(this,o,Be,Se,se,{})}}export{Ae as default};
