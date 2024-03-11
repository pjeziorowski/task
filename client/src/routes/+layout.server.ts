import { SERVER_URL } from "$env/static/private";
import favorites from "$lib/data";
import api from "$lib/server/api";
import type { LayoutServerLoad } from "./$types";

export const load = (async () => {
    const go_server_favorites = await api(SERVER_URL + "/favorites");
    if (!go_server_favorites.success) {
        console.error(
            "Failed to load favorites from server",
            go_server_favorites.error,
        );
    } else {
        console.log("Loaded favorites from server", go_server_favorites.data);
    }

    return {
        favorites: favorites,
    };
}) satisfies LayoutServerLoad;
