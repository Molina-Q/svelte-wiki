<script></script>

<div>
    <h2 class="title">Add a&nbsp;<span>Game</span></h2>
    <form action="">
        <label for="title">Title</label>
        <input type="text" id="title" name="title">
        
        <label for="description">Description</label>
        <textarea id="description" name="description"></textarea>
        
        <label for="image">Image</label>
        <input type="file" id="image" name="image">
        
        <label for="price">Price</label>
        <input type="number" id="price" name="price">
        
        <label for="platform">Platform</label>
        <input type="text" id="platform" name="platform">
        
        <button type="submit">Add Game</button>
    </form>
</div>

<style>
    form {
        width: 50vw;
        margin: 0 auto;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    label {
        font-size: 1.5rem;
    }

    input, textarea {
        padding: 0.5rem;
        font-size: 1.5rem;
    }

    button {
        padding: 0.5rem;
        font-size: 1.5rem;
        background-color: var(--svg-blue);
        color: var(--svg-white);
    }

    .title {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;

        font-size: 61px;
        color: var(--svg-blue);
    

        span {
            /* font-size: 67px; */
            color: var(--svg-orange);
            text-shadow: 5px 4px 0 var(--svg-blue);
        }
    }
</style>