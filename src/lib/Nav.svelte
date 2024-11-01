<script>
	import { Link } from "svelte-routing";
</script>

<nav class="navbar">
    <Link to="/">Home</Link>
    <Link to="/contact">About</Link>
</nav>

<style>
    .navbar {
        display: flex;
        justify-content: space-around;
        padding: 1em;
        background-color: #f1f1f1;
    }

    .navbar a {
        text-decoration: none;
        color: #333;
    }

    .navbar a:hover {
        color: #646cff;
    }
</style>