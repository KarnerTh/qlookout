import{w as r}from"./singletons.0e64877f.js";const l=()=>{let o;o=localStorage.notifications?JSON.parse(localStorage.notifications):null;const{subscribe:i,update:a,set:n}=r(o||[]),s=t=>{a(e=>[t,...e])},c=()=>{n([])};return i(t=>localStorage.notifications=JSON.stringify(t)),{subscribe:i,add:s,clear:c}},S=l();export{S as n};
