<script lang="ts">
    import { superForm } from "sveltekit-superforms/client";
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

    $: isFavorite = data.favorites
        .map((favorite) => favorite.animeId)
        .includes(data.anime.mal_id);
</script>

<div class="mb-8">
    <a class="button" href="/">Go back to list</a>
</div>

<div class="flex gap-8">
    <img
        src={data.anime.images.webp.image_url}
        alt={data.anime.title}
        class="h-96 w-64 rounded-xl"
    />
    <div>
        <h2>{data.anime.title}</h2>
        <p class="mb-4 flex-1">{data.anime.synopsis}</p>

        {#if !isFavorite}
            <form
                action="?/addToFavorites"
                method="post"
                use:addToFavoritesEnhance
            >
                <input
                    type="hidden"
                    name="animeId"
                    value={$addToFavoritesForm.animeId}
                />
                <input
                    type="hidden"
                    name="title"
                    value={$addToFavoritesForm.title}
                />
                <input
                    type="hidden"
                    name="image"
                    value={$addToFavoritesForm.image}
                />
                <button class="button" type="submit">Add to favorites</button>
            </form>
        {:else}
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
                <button class="button" type="submit">
                    Remove from favorites
                </button>
            </form>
        {/if}
    </div>
</div>

<style lang="postcss">
    .button {
        @apply block w-fit rounded-xl border border-gray-300 p-4 shadow transition-colors duration-200 ease-in-out hover:bg-zinc-50;
    }
</style>
