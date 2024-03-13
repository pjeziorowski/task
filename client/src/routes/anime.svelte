<script lang="ts">
    export let animeId: number;
    export let title: string;
    export let image: string;
</script>

<a
    href="/{animeId}"
    class="block h-full max-w-48 rounded-xl border border-gray-300 p-4 transition-transform duration-200 ease-in-out hover:scale-105"
>
    <h3 class="line-clamp-1 font-bold">{title}</h3>
    <img class="h-52 w-40 rounded-xl shadow" src={image} alt={title} />
</a>
