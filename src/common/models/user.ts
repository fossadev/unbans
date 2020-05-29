import {OAuthProvider} from "./provider";

export interface User {
  id: number;
  login: string;
  display_name: string;
  type: string;
  broadcaster_type: string;
  provider: OAuthProvider;
  provider_id: string;
  created_at: string;
  updated_at: string;
}
