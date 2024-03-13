<script lang="ts">
    export let animeId: number;
    export let title: string;
    export let image: string;
</script>

<h1 class="text-center font-semibold">{title}</h1>

<a href="/{animeId}">
    <img
        class="hover:cursor-pointer hover:opacity-50"
        src={image}
        alt={title}
        height="200px"
        width="200px"
    />
</a>
