import api from "$lib/server/api";
import { addFavorite, removeFavorite } from "$lib/server/favorites";
import { error, fail } from "@sveltejs/kit";
import { zod } from "sveltekit-superforms/adapters";
import { superValidate } from "sveltekit-superforms/server";
import { z } from "zod";
import type { Actions, PageServerLoad } from "./$types";

export type JikanAPIAnime = {
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
    const anime = await api<JikanAPIAnime>(`https://api.jikan.moe/v4/anime/${id}`);
    if (!anime.success) {
        console.error("Failed to fetch anime", anime.error);
        throw error(500, "Failed to fetch anime");
    }
    console.debug("anime", anime.data);
    return {
        removeFromFavoritesForm: await superValidate(
            {
                animeId: anime.data.data.mal_id,
            },
            zod(removeFromFavoritesSchema),
        ),
        addToFavoritesForm: await superValidate(
            {
                animeId: anime.data.data.mal_id,
                title: anime.data.data.title,
                image: anime.data.data.images.webp.image_url,
            },
            zod(addToFavoritesSchema),
        ),
        anime: anime.data.data,
    };
}) satisfies PageServerLoad;

const removeFromFavoritesSchema = z.object({
    animeId: z.number(),
});

const addToFavoritesSchema = z.object({
    animeId: z.number(),
    title: z.string(),
    image: z.string(),
});

export const actions = {
    addToFavorites: async ({ request, locals }) => {
        const form = await superValidate(request, zod(addToFavoritesSchema));

        if (!form.valid) {
            return fail(400, { form });
        }
        if (!locals.userId) {
            return fail(401, { message: "Unauthorized", form });
        }
        
        const inFavorites = await addFavorite(locals.userId, form.data)
        if (!inFavorites) {
            return fail(500, { message: "Failed to add favorite", form });
        }

        return { form };
    },
    removeFromFavorites: async ({ request, locals }) => {
        const form = await superValidate(
            request,
            zod(removeFromFavoritesSchema),
        );
        
        if (!form.valid) {
            return fail(400, { form });
        }
        if (!locals.userId) {
            return fail(401, { message: "Unauthorized", form });
        }

        const removed = await removeFavorite(locals.userId, form.data.animeId)
        if (!removed) {
            return fail(500, { message: "Failed to remove favorite", form });
        }

        return { form };
    },
} satisfies Actions;
