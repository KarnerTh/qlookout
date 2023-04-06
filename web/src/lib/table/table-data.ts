export interface TableDataText {
  type: "text";
  value: string;
}

export interface TableDataNumber {
  type: "number";
  value: number;
}

export interface TableDataBoolean {
  type: "boolean";
  value: boolean;
}

export type TableData = TableDataText | TableDataNumber | TableDataBoolean;

export interface TableRow {
  id: number;
  data: TableData[];
}
