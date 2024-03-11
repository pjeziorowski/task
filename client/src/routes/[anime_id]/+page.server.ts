import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import favorites from "$lib/data";

export type Anime = {
    data: {
        mal_id: number;
        title: string;
        url: string;
        synopsis: string;
        images: {
            webp: {
                image_url: string;
                small_image_url: string;
                large_image_url: string;
            };
        };
    };
};

export const load = (async ({ params }) => {
    const id = params.anime_id;
    const anime = await api<Anime>(
        `https://api.jikan.moe/v4/anime/${id}`,
    );
    if (!anime.success) {
        console.error("Failed to fetch anime", anime.error);
        throw error(500, "Failed to fetch anime");
    }
    console.debug("anime", anime.data);
    return {
        anime: anime.data.data,
    };
}) satisfies PageServerLoad;

export const actions = {
    addToFavorites: async ({ request }) => {
        const form = await request.formData();

        // validate form
        const mal_id = form.get("mal_id") as unknown as number;
        const title = form.get("title") as unknown as string;
        const image = form.get("image") as unknown as string;

        favorites.set(mal_id, { title: title, image: image });

        return { success: true };
    },
} satisfies Actions;
