"use client";

import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { createRecipe, listIngredients } from "@/services/nutrition/api";
import { generateObjectId } from "@/utils/bson_handling";

import '@/app/globals.css';
import { ApiListIngredientsResponse, ModelsIngredient } from "@/generated/services/nutrition/api";
import IngredientCard from "./ingredient_card";
import { useRouter } from "next/navigation";


interface FormData {
    title: string;
    description: string;
    ingredientIds: string[];
}

const CreateIngredientPage = () => {
    const [ingredientsState, updateIngredients] = useState<ModelsIngredient[]>();

    const router = useRouter();

    useEffect(() => {
        listIngredients().then((resp: ApiListIngredientsResponse) => {
            updateIngredients(resp.entities);
        });
    }, []);

    const [formData, setFormData] = useState<FormData>({
        title: '',
        description: '',
        ingredientIds: [],
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        await createRecipe({
            id: generateObjectId(),
            title: formData.title,
            description: formData.description,
            ingredientIds: formData.ingredientIds,
        });
        router.back();
    };

    const onIngredientSelect = (ingredient: ModelsIngredient) => {
        const currentIngredientIds = formData.ingredientIds;
        currentIngredientIds.push(ingredient.id);
        setFormData({
            ...formData,
            ingredientIds: currentIngredientIds,
        });
    }

    const onIngredientDeselect = (ingredient: ModelsIngredient) => {
        const currentIngredientIds = formData.ingredientIds;
        const ingredientIndex = currentIngredientIds.findIndex((id: string) => id === ingredient.id);
        if (ingredientIndex > -1) {
            currentIngredientIds.splice(ingredientIndex, 1);
        }
        setFormData({
            ...formData,
            ingredientIds: currentIngredientIds,
        });
    }

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="max-w-md mx-auto bg-white p-8 mt-10 rounded-lg shadow-lg">
                <h1 className="text-2xl font-bold mb-6 text-center">Create Ingredient</h1>
                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label htmlFor="title" className="block text-gray-700 font-bold mb-2">Title:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={formData.title}
                            onChange={handleChange}
                            required
                            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                        />
                    </div>
                    <div className="mb-4">
                        <label htmlFor="description" className="block text-gray-700 font-bold mb-2">Description:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={formData.description}
                            onChange={handleChange}
                            required
                            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 transition duration-300"
                    >
                        Create New Recipe
                    </button>
                </form>
            </div>
            <div className="min-h-screen bg-gray-100 py-10">
                <div className="container mx-auto px-4">
                    <h1 className="text-4xl font-bold text-center mb-8">Ingredients List</h1>
                    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
                        {ingredientsState?.map((ingredient: ModelsIngredient) => (
                            <IngredientCard key={ingredient.id} ingredient={ingredient} onSelect={onIngredientSelect} onDeselect={onIngredientDeselect} />
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default CreateIngredientPage;