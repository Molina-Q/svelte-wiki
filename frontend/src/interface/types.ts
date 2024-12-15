export interface IUser {
    id: number;
    lastName: string;
    firstName: string;
    email: string;
    password: string;
    role: string;
}

export interface IPlatform {
    id: number;
    name: string;
    description: string;
    image: string;
    releaseDate: string;
}

export interface IDeveloper {
    id: number;
    name: string;
    description: string;
    image: string;
    creationDate: string;
    games: IGame[];
}

export interface IGame {
    id: number;
    name: string;
    price: number;
    description: string;
    image: string;
    releaseDate: string;
    platform: IPlatform;
}

