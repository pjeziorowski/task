<script lang="ts">
    import { superForm } from "sveltekit-superforms/client";
    import Anime from "../anime.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    const addForm = superForm(data.addToFavoritesForm);
    const { form: addToFavoritesForm, enhance: addToFavoritesEnhance } =
        addForm;

    const removeForm = superForm(data.removeFromFavoritesForm);
    const {
        form: removeFromFavoritesForm,
        enhance: removeFromFavoritesEnhance,
    } = removeForm;
</script>

<Anime
    animeId={data.anime.mal_id}
    title={data.anime.title}
    image={data.anime.images.webp.image_url}
/>

<p>{data.anime.synopsis}</p>

<form action="?/addToFavorites" method="post" use:addToFavoritesEnhance>
    <input type="hidden" name="animeId" value={$addToFavoritesForm.animeId} />
    <input type="hidden" name="title" value={$addToFavoritesForm.title} />
    <input type="hidden" name="image" value={$addToFavoritesForm.image} />
    <button class="mt-4 rounded bg-gray-300 p-4" type="submit">
        Add to favorites
    </button>
</form>

<form
    action="?/removeFromFavorites"
    method="post"
    use:removeFromFavoritesEnhance
>
    <input
        type="hidden"
        name="animeId"
        value={$removeFromFavoritesForm.animeId}
    />
    <button class="mt-4 rounded bg-gray-300 p-4" type="submit">
        Remove from favorites
    </button>
</form>

<div class="mt-10">
    <a class="rounded bg-gray-300 p-4" href="/">Go back to list</a>
</div>
