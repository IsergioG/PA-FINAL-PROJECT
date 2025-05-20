export interface Kill {
    id: number;
    fullName: string;
    causeOfDeath?: string;
    details?: string;
    faceImageUrl: string | null;
    createdAt: string;
    // updatedAt: string;
}