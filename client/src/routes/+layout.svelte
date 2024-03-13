<script lang="ts">
    import { onNavigate } from "$app/navigation";
    import { page } from "$app/stores";
    import { flip } from "svelte/animate";
    import { quintOut } from "svelte/easing";
    import { fade } from "svelte/transition";
    import "../app.css";
    import type { LayoutData } from "./$types";
    import Anime from "./anime.svelte";

    const navbarStyle =
        "sticky top-0 z-10 bg-white/50 p-4 font-semibold backdrop-blur border-b border-gray-300 rounded-xl rounded-t-none shadow px-8 py-4";

    onNavigate((navigation) => {
        if (!document.startViewTransition) return;

        return new Promise((resolve) => {
            document.startViewTransition(async () => {
                resolve();
                await navigation.complete;
            });
        });
    });

    export let data: LayoutData;
</script>

<div class="max-w-7xl mx-auto flex min-h-screen">
    <div class="flex-1">
        <h2 class="{navbarStyle} rounded-r-none border-l">
            {$page.data["pageTitle"] || data.pageTitle}
        </h2>
        <div class="p-8">
            <slot />
        </div>
    </div>
    <div
        class="sticky top-0 h-screen w-64 flex-none items-center justify-center overflow-y-auto border-l border-gray-300 bg-white"
    >
        <h2 class="{navbarStyle} rounded-l-none border-r">
            Your favorites ({data.favorites.length})
        </h2>
        <div class="flex flex-col gap-8 p-8">
            {#each data.favorites as favorite (favorite.animeId)}
                <div
                    transition:fade
                    animate:flip={{
                        delay: 250,
                        duration: 250,
                        easing: quintOut,
                    }}
                >
                    <Anime
                        animeId={favorite.animeId}
                        title={favorite.title}
                        image={favorite.image}
                    />
                </div>
            {/each}
        </div>
    </div>
</div>
