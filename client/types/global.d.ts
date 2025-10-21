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
    img_url?: string | null;
    bio?: string | null;
    created_at?: string | null;
    updated_at?: string | null;
  }

  export interface Fyre {
    id: number;
    title: string;
    streak_count: number;
    user_id: number;
    bonfyre_id?: number | null;
    active_days: string;
    created_at?: string | null;
    updated_at?: string | null;
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
}
