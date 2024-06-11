"use client";

import RecipeForm from "./recipe_form";
import { createRecipe } from "@/services/nutrition/api";

const CreateIngredientPage = () => {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <RecipeForm onClick={createRecipe} />
        </div>
    );
};

export default CreateIngredientPage;