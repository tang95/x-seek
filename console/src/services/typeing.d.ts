declare namespace API {

    type AuthUser = {
        id: string;
        name: string;
        avatar: string;
        description?: string;
        role: string;
        token: string;
    }

    type User = {
        id: string;
        name: string;
        avatar: string;
        description?: string;
    }

    type Team = {
        id: string;
        name: string;
        description?: string;
        members?: User[];
    }

}
