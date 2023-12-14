<script>
    import Router from "svelte-spa-router";
    import IndexPage from "./pages/IndexPage.svelte";
    import NotFound from "./pages/NotFound.svelte";

    const routes = {
        "/": IndexPage,
        "*": NotFound,
    };
</script>

<div
    class="w-full min-h-screen text-black bg-neutral-100 dark:bg-neutral-950 dark:text-white"
>
    <Router {routes} />
</div>
