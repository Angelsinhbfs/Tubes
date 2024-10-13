import React from "react";
import '../styles/Tile.css'

interface PRT_TileProps {
    label: string;
    onClick?: () => void;
    imgURL?: string;
}
const PTile: React.FC<PRT_TileProps> = ({label, onClick}) => {
    return (
        <div className="tile-bg" onClick={onClick}>
            <div className="tile-overlay">
                <span className="overlay-text">{label}</span>
            </div>
        </div>
    );
};

export default PTile;