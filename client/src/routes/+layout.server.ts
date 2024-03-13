import { SERVER_URL } from "$env/static/private";
import api from "$lib/server/api";
import type { Favorite } from "$lib/server/favorites";
import { error } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load = (async ({ locals }) => {
    if (locals.userId) {
        const favorites = await api<Favorite[]>(
            SERVER_URL + `/users/${locals.userId}/favorites`,
        );

        if (!favorites.success) {
            throw error(500, "Failed to fetch user favorites");
        }

        return {
            favorites: favorites.data,
        };
    }

    return {
        favorites: [],
    };
}) satisfies LayoutServerLoad;
