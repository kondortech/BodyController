"use client";

import '@/app/globals.css';
import { ModelsIngredient } from '@/generated/services/nutrition/api';
import { createIngredient } from '@/services/nutrition/ingredients_api';
import { generateObjectId } from '@/utils/bson_handling';
import { ChangeEvent, FormEvent, useState } from 'react';


interface FormData {
    title: string;
    calories: string;
    proteins: string;
    carbs: string;
    fats: string;
}

export interface Props {
    onClick: (ingredient: ModelsIngredient) => void;
}

const IngredientForm: React.FC<Props> = ({ onClick }) => {
    const [formData, setFormData] = useState<FormData>({
        title: '',
        calories: '',
        proteins: '',
        carbs: '',
        fats: ''
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleSubmit = (e: FormEvent) => {
        e.preventDefault();
        console.log('Form data:', formData);
        onClick({
            id: generateObjectId(),
            title: formData.title,
            macrosNormalized: {
                calories: Number(formData.calories),
                proteins: Number(formData.proteins),
                carbs: Number(formData.carbs),
                fats: Number(formData.fats),
            }
        })
    };

    return (
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
                    <label htmlFor="calories" className="block text-gray-700 font-bold mb-2">Calories (per 100g):</label>
                    <input
                        type="number"
                        id="calories"
                        name="calories"
                        value={formData.calories}
                        onChange={handleChange}
                        required
                        className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                    />
                </div>
                <div className="mb-4">
                    <label htmlFor="proteins" className="block text-gray-700 font-bold mb-2">Proteins (g per 100g):</label>
                    <input
                        type="number"
                        id="proteins"
                        name="proteins"
                        value={formData.proteins}
                        onChange={handleChange}
                        required
                        className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                    />
                </div>
                <div className="mb-4">
                    <label htmlFor="carbs" className="block text-gray-700 font-bold mb-2">Carbs (g per 100g):</label>
                    <input
                        type="number"
                        id="carbs"
                        name="carbs"
                        value={formData.carbs}
                        onChange={handleChange}
                        required
                        className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                    />
                </div>
                <div className="mb-4">
                    <label htmlFor="fats" className="block text-gray-700 font-bold mb-2">Fats (g per 100g):</label>
                    <input
                        type="number"
                        id="fats"
                        name="fats"
                        value={formData.fats}
                        onChange={handleChange}
                        required
                        className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                    />
                </div>
                <button
                    type="submit"
                    className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 transition duration-300"
                >
                    Create Ingredient
                </button>
            </form>
        </div>
    );
};

export default IngredientForm;
