import { useMutation, useQueryClient } from "react-query"

export function useUsersQueries({ queryKey = 'users-list '} = {}) {
    const queryClient = useQueryClient()

    const updateMutation = useMutation(id => fetch(`/api/users/${id}`, {
        method: 'PUT',
        body: JSON.stringify({
            email: `New email ${Date.now()}`,
            name: 'New name',
        }),
    }), {
        onSuccess: () => queryClient.fetchQuery(queryKey),
    })

    const deleteMutation = useMutation(id => fetch(`/api/users/${id}`, {
        method: 'DELETE',
    }), {
        onSuccess: () => queryClient.fetchQuery(queryKey),
    })
    
    const createMutation = useMutation(user => fetch(`/api/users`, {
        method: 'POST',
        body: JSON.stringify(user),
    }), {
        onSuccess: () => queryClient.fetchQuery(queryKey),
    })

    return {
        updateMutation,
        deleteMutation,
        createMutation,
    }
}

export async function fetchUsers() {
    const res = await fetch('/api/users')
    const data = await res.json()
    return data
}

export async function fetchUserById(id) {
    const res = await fetch(`/api/users/${id}`)
    const data = await res.json()
    return data
}