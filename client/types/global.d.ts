/* Place any types that are used globally, these types do not need to be imported to use */

/**
 * All database tables are defined here
 * if you need to use a modified form, make a new interface in the local directory
 */
declare namespace Models {
  export interface User {
    id: number;
    name: string;
    email: string;
    imgURL?: string;
    bio?: string;
    created_at?: string | null;
    updated_at?: string | null;
  }

  export interface Fyre {
    id: number;
    title: string;
    userID: number;
    streakCount: number | 0;
    bonfyreID?: number | null;
    created_at?: string | null;
    updated_at?: string | null;
  }

  export interface Goal {
    fyreID: number;
    description: string;
    milestone: number;
  }

  export interface Bonfyre {
    id: number;
  }

  export interface Friend {
    userID1: number;
    userID2: number;
  }
}
