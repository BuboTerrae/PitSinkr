export interface ThemeType {
  theme: string;
  toggleTheme: () => void;
}

export type LogType = {
  client_ip:string;
  domain:string;
  id: number;
  query_type:string;
  timestamp:Date|string;
};
