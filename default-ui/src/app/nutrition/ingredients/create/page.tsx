"use client";

import { ModelsIngredient } from "@/generated/services/nutrition/api";
import IngredientForm from "./ingredient_form";

const onSubmit = (_: ModelsIngredient) => { }

const CreateIngredientPage = () => {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <IngredientForm />
        </div>
    );
};

export default CreateIngredientPage;