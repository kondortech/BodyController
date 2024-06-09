"use client";

import IngredientForm from "./ingredient_form";
import { createIngredient } from "@/services/nutrition/ingredients_api";

const CreateIngredientPage = () => {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <IngredientForm onClick={createIngredient} />
        </div>
    );
};

export default CreateIngredientPage;