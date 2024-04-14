import React from 'react';
import styles from './styles.module.css'

type Props = {
    link: string,
    title: string,
    description?: string,
};

export const Tile = (props: Props): JSX.Element => {
    return (
        <a href={props.link} className={styles.tile}>
            <div className={styles.tile}>
                <div className={styles.tile_content}>
                    <p className={styles.tile_title}>{props.title}</p>
                    {(props.description !== undefined) ? <div className={styles.tile_description}>{props.description}</div> : undefined}
                </div>
            </div>
        </a>
    );
}