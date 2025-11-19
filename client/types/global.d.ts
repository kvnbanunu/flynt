/* Place any types that are used globally, these types do not need to be imported to use */

/**
 * All database tables are defined here
 * if you need to use a modified form, make a new interface in the local directory
 */
declare namespace Models {
  export interface User {
    id: number;
    username: string;
    name: string;
    email: string;
    img_url?: string;
    bio?: string;
    timezone: string;
    fyre_total: number;
    created_at?: string;
    updated_at?: string;
  }

  export interface Fyre {
    id: number;
    title: string;
    streak_count: number;
    user_id: number;
    bonfyre_id?: number;
    active_days: string;
    is_checked: boolean;
    last_checked_at?: string;
    last_checked_at_prev?: string;
    created_at?: string;
    updated_at?: string;
  }

  export interface Goal {
    fyre_id: number;
    description: string;
    milestone: number;
  }

  export interface Bonfyre {
    id: number;
  }

  export interface Friend {
    user_id_1: number;
    user_id_2: number;
  }

  export interface SocialPost {
    id: number;
    user_id: number;
    fyre_id: number;
    type: "dailycheck" | "milestone";
    content: string;
  }
}
