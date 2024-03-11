<script lang="ts">
    import type { PageData } from "./$types";
    import Anime from "./anime.svelte";

    export let data: PageData;
</script>

{#each data.recommended.data as recommendation}
    <div class="mb-4 grid grid-cols-2 rounded border border-gray-800">
        <h1 class="col-span-2 border-b border-gray-800 p-2">
            {recommendation.content}
        </h1>
        {#each recommendation.entry as subRecommendation}
            <div class="mb-2 p-2">
                <Anime
                    title={subRecommendation.title}
                    mal_id={subRecommendation.mal_id}
                    image={subRecommendation.images.webp.image_url}
                />
            </div>
        {/each}
    </div>
{/each}
