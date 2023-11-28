import{S as te,i as ne,s as ae,w as L,R as U,y as T,a as b,k as Q,z as v,c as I,l as j,m as G,h as B,n as M,A as y,b as W,D as g,N as re,Q as oe,T as X,g as S,d as k,B as x,F as le}from"../chunks/index.f49fd5c2.js";import{P as ue,B as se,g as Y}from"../chunks/PageHeader.b42d5f8a.js";import{p as me}from"../chunks/stores.6683fd3d.js";import{F as Z,R as ie}from"../chunks/RuleInputs.63f4425d.js";import{F as ee}from"../chunks/FormInputField.6e32f34e.js";import{k as pe,m as fe}from"../chunks/svelte-apollo.e94973ad.js";const de=pe`
  mutation CreateRule(
    $lookoutId: Int!
    $columnName: String!
    $columnType: String!
    $rowIndex: Int!
    $exactValue: String
    $lessThan: String
    $greaterThan: String
    $shouldBeNull: Boolean
  ) {
    createRule(
      data: {
        lookoutId: $lookoutId
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
`,ce=()=>fe(de);function ge(r){let s,f,o,t,l,h,d,C,m,F,V,n,i,w,p,N,_,q,$,c,R,O,D;s=new ue({props:{title:"Create Rule",backAction:r[5]}}),l=new ee({props:{name:"columnName",required:!0,label:"Column name",placeholder:"columnName"}}),d=new ee({props:{name:"rowIndex",required:!0,type:"number",label:"Row index",placeholder:"0"}});function z(e){r[6](e)}let E={name:"columnType",required:!0,label:"Column Type",options:[{name:"text",title:"text"},{name:"int",title:"int"},{name:"float",title:"float"}]};r[0]!==void 0&&(E.value=r[0]),m=new Z({props:E}),L.push(()=>U(m,"value",z));function H(e){r[7](e)}let P={name:"ruleType",required:!0,label:"Rule Type",options:[{name:"exact",title:"exact value"},{name:"null",title:"should be null"},{name:"less",title:"less than"},{name:"greater",title:"greater than"},{name:"between",title:"between"}]};return r[1]!==void 0&&(P.value=r[1]),n=new Z({props:P}),L.push(()=>U(n,"value",H)),p=new ie({props:{ruleType:r[1],inputType:r[2]}}),c=new se({props:{title:"Create",type:"submit",leadingIcon:"check"}}),{c(){T(s.$$.fragment),f=b(),o=Q("form"),t=Q("div"),T(l.$$.fragment),h=b(),T(d.$$.fragment),C=b(),T(m.$$.fragment),V=b(),T(n.$$.fragment),w=b(),T(p.$$.fragment),N=b(),_=Q("div"),q=b(),$=Q("div"),T(c.$$.fragment),this.h()},l(e){v(s.$$.fragment,e),f=I(e),o=j(e,"FORM",{method:!0,class:!0});var u=G(o);t=j(u,"DIV",{class:!0});var a=G(t);v(l.$$.fragment,a),h=I(a),v(d.$$.fragment,a),C=I(a),v(m.$$.fragment,a),V=I(a),v(n.$$.fragment,a),w=I(a),v(p.$$.fragment,a),N=I(a),_=j(a,"DIV",{class:!0}),G(_).forEach(B),q=I(a),$=j(a,"DIV",{class:!0});var A=G($);v(c.$$.fragment,A),A.forEach(B),a.forEach(B),u.forEach(B),this.h()},h(){M(_,"class","w-full"),M($,"class","mt-4 mx-3 w-full flex flex-row-reverse"),M(t,"class","flex flex-wrap -mx-3 mb-6"),M(o,"method","POST"),M(o,"class","w-full max-w-lg")},m(e,u){y(s,e,u),W(e,f,u),W(e,o,u),g(o,t),y(l,t,null),g(t,h),y(d,t,null),g(t,C),y(m,t,null),g(t,V),y(n,t,null),g(t,w),y(p,t,null),g(t,N),g(t,_),g(t,q),g(t,$),y(c,$,null),R=!0,O||(D=re(o,"submit",oe(r[4])),O=!0)},p(e,[u]){const a={};u&8&&(a.backAction=e[5]),s.$set(a);const A={};!F&&u&1&&(F=!0,A.value=e[0],X(()=>F=!1)),m.$set(A);const K={};!i&&u&2&&(i=!0,K.value=e[1],X(()=>i=!1)),n.$set(K);const J={};u&2&&(J.ruleType=e[1]),u&4&&(J.inputType=e[2]),p.$set(J)},i(e){R||(S(s.$$.fragment,e),S(l.$$.fragment,e),S(d.$$.fragment,e),S(m.$$.fragment,e),S(n.$$.fragment,e),S(p.$$.fragment,e),S(c.$$.fragment,e),R=!0)},o(e){k(s.$$.fragment,e),k(l.$$.fragment,e),k(d.$$.fragment,e),k(m.$$.fragment,e),k(n.$$.fragment,e),k(p.$$.fragment,e),k(c.$$.fragment,e),R=!1},d(e){x(s,e),e&&B(f),e&&B(o),x(l),x(d),x(m),x(n),x(p),x(c),O=!1,D()}}}function $e(r,s,f){let o;le(r,me,n=>f(3,o=n));let t,l,h;const d=ce(),C=async n=>{var D,z,E,H,P;const i=new FormData(n.target),w=+o.params.lookoutId,p=(D=i.get("columnName"))==null?void 0:D.toString(),N=(z=i.get("columnType"))==null?void 0:z.toString(),_=+(i.get("rowIndex")??0),q=(E=i.get("exactValue"))==null?void 0:E.toString(),$=(H=i.get("lessThan"))==null?void 0:H.toString(),c=(P=i.get("greaterThan"))==null?void 0:P.toString(),R=i.has("shouldBeNull");if(!w||!p||!N){alert("Something went wrong");return}if((await d({variables:{lookoutId:w,columnName:p,columnType:N,rowIndex:_,exactValue:q,lessThan:$,greaterThan:c,shouldBeNull:R}})).errors){alert("Something went wrong");return}Y(`/lookout/${w}`,{state:{refetch:!0}})},m=()=>Y(`/lookout/${o.params.lookoutId}`);function F(n){l=n,f(0,l)}function V(n){t=n,f(1,t)}return r.$$.update=()=>{r.$$.dirty&1&&(["int","float"].includes(l)?f(2,h="number"):f(2,h="text"))},[l,t,h,o,C,m,F,V]}class Ie extends te{constructor(s){super(),ne(this,s,$e,ge,ae,{})}}export{Ie as default};
