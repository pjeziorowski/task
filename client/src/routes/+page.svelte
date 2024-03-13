<script lang="ts">
    import type { PageData } from "./$types";
    import Anime from "./anime.svelte";

    export let data: PageData;
</script>

<article class="flex flex-col gap-8">
    {#each data.recommended.data as recommendation}
        <section
            class="flex flex-col gap-8 rounded-xl border border-gray-300 bg-white p-4 md:flex-row"
        >
            <div class="flex gap-4">
                {#each recommendation.entry as subRecommendation}
                    <Anime
                        animeId={subRecommendation.mal_id}
                        title={subRecommendation.title}
                        image={subRecommendation.images.webp.image_url}
                    />
                {/each}
            </div>
            <div class="flex-1">
                <p class="line-clamp-[10]">
                    {recommendation.content}
                </p>
            </div>
        </section>
    {/each}
</article>
