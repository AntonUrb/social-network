import React from 'react';
import GroupInformation from './GroupInformation';
import JoinRequestsButton from '../buttons/JoinRequestsButton';
import InviteGroupButton from '../buttons/InviteGroupButton';
import Post from '../postcreation/Post';

interface GroupPageInfoProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const GroupPageInfo: React.FC<GroupPageInfoProps> = ({ title, text, pictureUrl }) => {
    return (
        <div>
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
                {/* Group Info*/}
                <GroupInformation
                    title={title} // Pass title prop to GroupContent
                    text={text}
                    pictureUrl={pictureUrl}
                    placeholderTitle="Shoe Emporium"
                    placeholderText="Join us for footwear everything!"
                    placeholderPictureUrl="https://iili.io/J1ucEoF.jpg"
                />
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between', border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
                {/* Invite People */}
                <InviteGroupButton/>
                <div style={{ width: '20px' }}></div>
                {/* Requests */}
                <JoinRequestsButton/>
            </div>
            
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px' }}>
            <h3 style={{ color: 'white', fontWeight:'bold', fontSize: '20px'}}>People in Group</h3>
            </div>
            
            {/* People in group list */}
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', height: '50vh', padding: '20px', marginBottom: '20px', overflowY: 'auto' }}>
                {/* List */}
                <ul style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    {/* Map through the list of people and render each item */}
                    
                </ul>
            </div>
        </div>
    );
};

export default GroupPageInfo;
