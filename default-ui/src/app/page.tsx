import { Tile } from "@/components/tile";
import { Collapsible } from "@/components/collapsible";
import styles from './styles.module.css'

export default function Home() {
	return (
		<main className="flex min-h-screen flex-col items-center p-24">
			<Collapsible title="Nutrition">
				<div className={styles.section}>
					<Tile link="/nutrition/ingredients" title="Ingredients" description="Fresh ingredients with verified macros" />
					<Tile link="/nutrition/recipes" title="Recipes" description="Delicious recipes with verified examples" />
					<Tile link="/nutrition/planning" title="Planning" description="Your nutrition plan" />
				</div>
			</Collapsible>
		</main>
	);
}
