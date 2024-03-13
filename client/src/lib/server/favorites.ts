import { SERVER_URL } from "$env/static/private";
import api from "./api";

export type Favorite = {
    animeId: number;
    userId: string;
    title: string;
    image: string;
};

export async function getFavorites(userId: string) {
    return await api<Favorite[]>(`${SERVER_URL}/users/${userId}/favorites`);
}

export async function addFavorite(
    userId: string,
    anime: {
        animeId: number;
        title: string;
        image: string;
    },
) {
    return await api(`${SERVER_URL}/users/${userId}/favorites`, {
        method: "PUT",
        body: anime,
    });
}

export async function removeFavorite(userId: string, animeId: number) {
    return await api(`${SERVER_URL}/users/${userId}/favorites/${animeId}`, {
        method: "DELETE",
    });
}
