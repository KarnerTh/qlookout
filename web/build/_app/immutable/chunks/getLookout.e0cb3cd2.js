import{k as e,q as u}from"./svelte-apollo.d2e69eac.js";const l=e`
  query Lookout($id: Int!) {
    lookout(id: $id) {
      id
      name
      cron
      query
      notifyLocal
      notifyMail
      rules {
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
  }
`,t=o=>u(l,{variables:{id:o}});export{t as u};
