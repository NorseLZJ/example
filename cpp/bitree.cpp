#include <stdio.h>
#include <stdlib.h>
#include <sys/malloc.h>

typedef struct Node
{
    int data;
    struct Node *lchild, *rchild;
} BiTreeNode, *BiTree;

BiTree BSTSearch(BiTree T, int x);
int BSTInsert(BiTree *T, int x);
void InOrderTraveres(BiTree T);

int main(int argc, char **argv)
{
    BiTree T = NULL, p;
    int table[] = {55, 33, 44, 66, 99, 77, 88, 22, 11};
    int n = sizeof(table) / sizeof(table[0]);
    int x, i;
    for (i = 0; i < n; i++)
        BSTInsert(&T, table[i]);
    
    printf("\n---------\n");
    InOrderTraveres(T);
    for(;;){
        
        printf("\ninput search num:");
        scanf("%d", &x);
        p = BSTSearch(T, x);
        if (p != NULL)
            printf("\nfind it's :%d", x);
        else
            printf("\ncan't find it's :%d", x);
    }
    
    return 0;
}

BiTree BSTSearch(BiTree T, int x)
{
    BiTreeNode *p;
    if (T != NULL)
    {
        p = T;
        while (p != NULL)
        {
            if (p->data == x)
                return p;
            else if (x < p->data)
                p = p->lchild;
            else
                p = p->rchild;
        }
    }
    return NULL;
}

int BSTInsert(BiTree *T, int x)
{
    BiTreeNode *p, *cur, *parent = NULL;
    cur = *T;
    while (cur != NULL)
    {
        if (cur->data == x)
            return 0;
        
        parent = cur;
        if (x < cur->data)
            cur = cur->lchild;
        else
            cur = cur->rchild;
    }
    p = (BiTreeNode *)malloc(sizeof(BiTreeNode));
    if (!p)
        exit(1);
    
    p->data = x;
    p->lchild = NULL;
    p->rchild = NULL;
    if (!parent)
        *T = p;
    else if (x < parent->data)
        parent->lchild = p;
    else
        parent->rchild = p;
    return 1;
}

void InOrderTraveres(BiTree T)
{
    if (!T)
        return;
    
    InOrderTraveres(T->lchild);
    printf("%4d", T->data);
    InOrderTraveres(T->rchild);
}
