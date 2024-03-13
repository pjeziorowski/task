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

    const buttonClass = "rounded-xl p-4 border border-gray-300 shadow";

    $: isFavorite = data.favorites
        .map((favorite) => favorite.animeId)
        .includes(data.anime.mal_id);
</script>

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
                <button class={buttonClass} type="submit">
                    Add to favorites
                </button>
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
                <button class={buttonClass} type="submit">
                    Remove from favorites
                </button>
            </form>
        {/if}
    </div>
</div>

<div class="mt-8">
    <a class={buttonClass} href="/">Go back to list</a>
</div>
