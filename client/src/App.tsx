import React, {useState} from 'react';
import Modal from './components/Modal';
import PTile from './components/PRT_Tile';
import UploadModal from "./components/UploadModal";
import "./styles/Main.css"

const App: React.FC = () => {
    const [isModalOpen, setModalOpen] = useState(false);
    const [isUploadOpen, setUploadOpen] = useState(false);

    const openModal = () => setModalOpen(true);
    const closeModal = () => setModalOpen(false);
    const openUpload = () => setUploadOpen(true);
    const closeUpload = () => setUploadOpen(false);

    return (
        <div>
            <PTile label="Test Tile" onClick={openModal}/>
            <Modal isOpen={isModalOpen} onClose={closeModal}>
                <h2>Modal Title</h2>
                <p>This is the modal content.</p>
            </Modal>
            <PTile label="Upload" onClick={openUpload}/>
            <UploadModal isOpen={isUploadOpen} onRequestClose={closeUpload}/>
        </div>
    );
};

export default App;