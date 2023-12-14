<script>
    const promise = fetch("/api/files")
        .then((res) => {
            return res.json();
        })
        .catch((err) => {
            console.error(err);
            throw new Error("Failed to fetch files.");
        });
</script>

<div class="flex flex-col gap-4 p-6">
    <form
        action="/api/upload"
        method="post"
        enctype="multipart/form-data"
        id="form"
        on:submit={async (e) => {
            e.preventDefault();
            const formData = new FormData(e.currentTarget);
            const res = await fetch("/api/upload", {
                method: "POST",
                body: formData,
            });
            if (res.ok) {
                console.log("Uploaded file.");
            } else {
                console.error("Failed to upload file.");
            }
        }}
    >
        <input
            type="file"
            name="file"
            id="fileInput"
            class="hidden"
            accept="image/*"
            on:change={() => {
                const form = document.getElementById("form");
                form.submit();
            }}
        />
    </form>

    <button
        class="px-2 py-1 text-white bg-black rounded-md dark:bg-white dark:text-black w-fit"
        id="uploadButton"
        on:click={() => {
            const fileInput = document.getElementById("fileInput");
            fileInput.click();
        }}
    >
        Upload
    </button>

    <div class="flex flex-col gap-2">
        {#await promise then files}
            {#each files as file}
                <a href={`/api/files/${file}`}>{file}</a>
            {/each}
        {:catch error}
            <p style="color: red">{error.message}</p>
        {/await}
    </div>
</div>
