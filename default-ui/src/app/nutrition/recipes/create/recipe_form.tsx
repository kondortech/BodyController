"use client";

import '@/app/globals.css';
import { ModelsRecipe } from '@/generated/services/nutrition/api';
import { generateObjectId } from '@/utils/bson_handling';
import { ChangeEvent, FormEvent, useState } from 'react';


interface FormData {
    title: string;
    description: string;
}

export interface Props {
    onClick: (instance: ModelsRecipe) => void;
}

const RecipeForm: React.FC<Props> = ({ onClick }) => {
    const [formData, setFormData] = useState<FormData>({
        title: '',
        description: '',
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
            description: formData.description,
            ingredientIds: [],
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
    );
};

export default RecipeForm;
